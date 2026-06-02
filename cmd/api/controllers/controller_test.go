package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"restapis-go/internal/middleware"
	"restapis-go/internal/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	user "restapis-go/internal/entities"
)

func init() {
	gin.SetMode(gin.TestMode)
	os.Setenv("JWT_SECRET", "test-secret-key")
}

// setupTestDB cria um banco SQLite em memória para os testes
func setupTestDB(t *testing.T) {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&user.User{}); err != nil {
		t.Fatalf("failed to migrate test db: %v", err)
	}
	repository.Database = db
}

// validToken gera um token JWT válido para autenticação nos testes
func validToken() string {
	claims := middleware.Claims{
		UserID: 1,
		User:   "TestUser",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("test-secret-key"))
	return tokenString
}

func setupRouter() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/api/v1")
	v1.POST("/login", Login)

	protected := v1.Group("/")
	protected.Use(middleware.Auth())
	{
		protected.GET("/", ListUsers)
		protected.GET("/:ID", ListUserByID)
		protected.POST("/post", NewUser)
		protected.DELETE("/delete/:ID", DeleteUser)
		protected.PATCH("/update/:ID", UpdateUser)
	}
	return r
}

// ---- NewUser ----

func TestNewUser_Sucesso(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	body, _ := json.Marshal(map[string]string{"name": "Arthur"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/post", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("status = %d, want %d | body: %s", w.Code, http.StatusCreated, w.Body.String())
	}
}

func TestNewUser_NomeFaltando(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	body, _ := json.Marshal(map[string]string{})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/post", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestNewUser_NomeInvalido(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	body, _ := json.Marshal(map[string]string{"name": "Arthur123"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/post", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestNewUser_SemAuth(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	body, _ := json.Marshal(map[string]string{"name": "Arthur"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/post", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", w.Code, http.StatusUnauthorized)
	}
}

func TestNewUser_BodyInvalido(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/post", bytes.NewBufferString("not-json"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

// ---- ListUsers ----

func TestListUsers_Vazio(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/", nil)
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestListUsers_ComUsuarios(t *testing.T) {
	setupTestDB(t)
	repository.Database.Create(&user.User{Name: "Arthur"})
	repository.Database.Create(&user.User{Name: "Maria"})

	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/", nil)
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}

	var users []user.User
	if err := json.Unmarshal(w.Body.Bytes(), &users); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}

// ---- ListUserByID ----

func TestListUserByID_Encontrado(t *testing.T) {
	setupTestDB(t)
	u := user.User{Name: "Arthur"}
	repository.Database.Create(&u)

	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/1", nil)
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d | body: %s", w.Code, http.StatusOK, w.Body.String())
	}
}

func TestListUserByID_NaoEncontrado(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/999", nil)
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

// ---- DeleteUser ----

func TestDeleteUser_Sucesso(t *testing.T) {
	setupTestDB(t)
	u := user.User{Name: "Arthur"}
	repository.Database.Create(&u)

	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/delete/1", nil)
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d | body: %s", w.Code, http.StatusOK, w.Body.String())
	}
}

func TestDeleteUser_NaoEncontrado(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/delete/999", nil)
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

// ---- UpdateUser ----

func TestUpdateUser_Sucesso(t *testing.T) {
	setupTestDB(t)
	u := user.User{Name: "Arthur"}
	repository.Database.Create(&u)

	r := setupRouter()
	body, _ := json.Marshal(map[string]string{"name": "ArthurAtualizado"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/api/v1/update/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d | body: %s", w.Code, http.StatusOK, w.Body.String())
	}
}

func TestUpdateUser_NomeFaltando(t *testing.T) {
	setupTestDB(t)
	u := user.User{Name: "Arthur"}
	repository.Database.Create(&u)

	r := setupRouter()
	body, _ := json.Marshal(map[string]string{})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/api/v1/update/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestUpdateUser_NomeInvalido(t *testing.T) {
	setupTestDB(t)
	u := user.User{Name: "Arthur"}
	repository.Database.Create(&u)

	r := setupRouter()
	body, _ := json.Marshal(map[string]string{"name": "123Invalid"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/api/v1/update/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+validToken())

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

// ---- Login ----

func TestLogin_Sucesso(t *testing.T) {
	setupTestDB(t)
	repository.Database.Create(&user.User{Name: "Arthur"})

	r := setupRouter()
	body, _ := json.Marshal(map[string]string{"name": "Arthur"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d | body: %s", w.Code, http.StatusOK, w.Body.String())
	}

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["Token"] == "" {
		t.Error("expected non-empty Token in response")
	}
}

func TestLogin_UsuarioNaoExiste(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	body, _ := json.Marshal(map[string]string{"name": "NaoExiste"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestLogin_NomeFaltando(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	body, _ := json.Marshal(map[string]string{})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestLogin_BodyInvalido(t *testing.T) {
	setupTestDB(t)
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBufferString("not-json"))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}