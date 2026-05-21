# 🚀 RESTAPIs-go

A REST API built with Go, Gin, GORM and PostgreSQL (Neon), following a clean architecture with use-cases, repository pattern and Swagger documentation.

## 🛠️ Tech Stack

- [Go](https://golang.org/)
- [Gin](https://gin-gonic.com/) — HTTP framework
- [GORM](https://gorm.io/) — ORM
- [PostgreSQL (Neon)](https://neon.tech/) — Database
- [Swagger](https://swagger.io/) — API documentation
- [Docker](https://www.docker.com/)

## 📋 Prerequisites

- [Go 1.21+](https://golang.org/dl/)
- [Make](https://www.gnu.org/software/make/)

## ⚙️ Setup

1. Clone the repository
```bash
git clone https://github.com/ArthurLoffi/RESTAPIs-go.git
cd RESTAPIs-go
```

2. Create a `.env` file at the root of the project
```env
DATABASE_URL=postgresql://<user>:<password>@<host>/<database>?sslmode=require
```

3. Install dependencies
```bash
go mod download
```

## 🚀 Running

**With Make:**
```bash
make build   # compile
make run     # run binary
make dev     # run without compiling
```

**Without Make:**
```bash
go run ./cmd/api/
```

**Binary:**
```bash
./api.exe
```

## 📡 Endpoints

| Method | Route | Description |
|--------|-------|-------------|
| GET | `/api/v1/users` | List all users |
| POST | `/api/v1/users/:name` | Create user |
| DELETE | `/api/v1/users/:id` | Delete user |

## 📄 Swagger

After running the project, access:
```
http://localhost:8080/swagger/index.html
```

To regenerate docs:
```bash
make docs
```

## 📁 Project Structure

```
├── cmd/
│   └── api/
│       ├── main.go
│       ├── routes.go
│       └── controllers/
├── docs/
├── internal/
│   ├── entities/
│   │   └── user.go
│   ├── repository/
│   │   └── repository.go
│   └── use-cases/
│       ├── create-user.go
│       ├── delete-user.go
│       └── get-user.go
├── .env
├── Dockerfile
├── go.mod
└── MakeFile
```