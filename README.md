# RESTAPIs-go

REST API built with **Go**, **Gin**, **GORM** and **PostgreSQL (Neon)**, following Clean Architecture principles with use-cases, repository pattern, JWT authentication, rate limiting, request logging and Swagger documentation.

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | [Go 1.26+](https://golang.org/) |
| HTTP Framework | [Gin](https://gin-gonic.com/) |
| ORM | [GORM](https://gorm.io/) |
| Database | [PostgreSQL via Neon](https://neon.tech/) |
| Authentication | [JWT (golang-jwt/jwt v5)](https://github.com/golang-jwt/jwt) |
| Password Hashing | [bcrypt (golang.org/x/crypto)](https://pkg.go.dev/golang.org/x/crypto/bcrypt) |
| API Docs | [Swagger (swaggo)](https://github.com/swaggo/swag) |
| Containerization | [Docker](https://www.docker.com/) |
| Rate Limiting | [golang.org/x/time](https://pkg.go.dev/golang.org/x/time/rate) |

---

## Project Structure

```
├── cmd/
│   └── api/
│       ├── main.go               # Entry point, server bootstrap
│       ├── routes.go             # Route registration
│       └── controllers/
│           ├── auth.go           # Login handler
│           └── user-controller.go # User CRUD handlers
├── docs/                         # Swagger generated files
├── internal/
│   ├── entities/
│   │   └── user.go               # User model (GORM)
│   ├── error/
│   │   └── error.go              # Centralized error type and responder
│   ├── middleware/
│   │   ├── auth.go               # JWT validation middleware
│   │   └── logger.go             # Request logger middleware
│   ├── repository/
│   │   └── repository.go         # Database connection and AutoMigrate
│   └── use-cases/
│       ├── auth-user.go          # Login: name lookup + bcrypt verify
│       ├── create-user.go        # Create user with bcrypt hash
│       ├── get-user.go           # List users (paginated) + get by ID
│       ├── update-user.go        # Update user name
│       └── delete-user.go        # Delete user by ID
├── pkg/
│   └── validateName.go           # Name validation (letters + accents only)
├── .dockerignore
├── .gitignore
├── Dockerfile
├── Makefile
├── go.mod
└── go.sum
```

---

## Prerequisites

- Go 1.26+
- Make
- A [Neon](https://neon.tech/) PostgreSQL database (or any PostgreSQL instance)

---

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/ArthurLoffi/RESTAPIs-go.git
cd RESTAPIs-go
```

### 2. Create a `.env` file at the project root

```env
DATABASE_URL=postgresql://<user>:<password>@<host>/<database>?sslmode=require
JWT_SECRET=your-secret-key
GIN_MODE=debug
```

### 3. Install dependencies

```bash
make init
```

---

## Running

**With Make:**

```bash
make build   # compiles the binary
make run     # runs the compiled binary
```

**Without Make:**

```bash
go run ./cmd/api/
```

**With Docker:**

```bash
docker build -t restapis-go .
docker run -p 8080:8080 --env-file .env restapis-go
```

---

## API Endpoints

Base path: `/api/v1`

| Method | Route | Auth Required | Description |
|--------|-------|:---:|---|
| `GET` | `/healthy` | ✗ | Health check |
| `POST` | `/login` | ✗ | Login — returns a JWT token |
| `GET` | `/users?page=1&pageSize=10` | ✓ | List all users (paginated) |
| `GET` | `/user/:ID` | ✓ | Get user by ID |
| `POST` | `/post` | ✓ | Create a new user |
| `PATCH` | `/update/:ID` | ✓ | Update user name |
| `DELETE` | `/delete/:ID` | ✓ | Delete user by ID |

Protected routes require the `Authorization` header:

```
Authorization: Bearer <token>
```

---

## Authentication Flow

### Register a user

```bash
POST /api/v1/post
Content-Type: application/json
Authorization: Bearer <token>

{
  "Name": "Arthur",
  "Password": "yourpassword"
}
```

Passwords are **hashed with bcrypt** before being stored. Plain-text passwords are never persisted.

### Login

```bash
POST /api/v1/login
Content-Type: application/json

{
  "name": "Arthur",
  "Password": "yourpassword"
}
```

**Response:**

```json
{
  "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

The token encodes `user_id` and `name` in the claims and expires after **24 hours**.

> ⚠️ Both "user not found" and "wrong password" return the same `401 Invalid credentials` response to prevent user enumeration attacks.

---

## Name Validation

User names are validated against the following rule:

- Only letters (including accented characters like `ç`, `ã`, `é`) and spaces are accepted.
- Numbers and special characters are rejected.

This is enforced by `pkg/validateName.go` using a compiled regex.

---

## Middleware

| Middleware | Description |
|---|---|
| `Logger` | Logs method, path, status code and request duration for every request |
| `Auth` | Validates the `Bearer` JWT token and injects `user_id` and `name` into the Gin context |
| `Limiter` | Rate limiting applied to all `/api/v1` routes |

---

## Swagger Docs

After running the project, access the interactive docs at:

```
http://localhost:8080/api/v1/swagger/index.html
```

To regenerate the docs after changing Swagger annotations:

```bash
make docs
```

---

## Tests

Tests use an **in-memory SQLite** database and do not require a running PostgreSQL instance.

```bash
go test ./... -v
```

Check coverage:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

> The test setup script is defined in `Makefile` as `make test`.

---

## Error Format

All errors follow a consistent JSON format:

```json
{
  "status": 400,
  "message": "invalid or missing name"
}
```

---

## License

This project is for educational purposes.