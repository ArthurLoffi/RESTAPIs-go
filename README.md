# RESTAPIs-go

REST API built with Go, Gin, GORM and PostgreSQL, following Clean Architecture with use-cases, repository pattern, JWT authentication and Swagger documentation.

## Tech Stack

- [Go](https://golang.org/)
- [Gin](https://gin-gonic.com/) — HTTP framework
- [GORM](https://gorm.io/) — ORM
- [PostgreSQL (Neon)](https://neon.tech/) — Database
- [JWT](https://github.com/golang-jwt/jwt) — Authentication
- [Swagger](https://swagger.io/) — API documentation
- [Docker](https://www.docker.com/)

## Prerequisites

- Go 1.22+
- Make

## Setup

1. Clone the repository

```bash
git clone https://github.com/ArthurLoffi/RESTAPIs-go.git
cd RESTAPIs-go
```

2. Create a `.env` file at the root of the project

```env
DATABASE_URL=postgresql://<user>:<password>@<host>/<database>?sslmode=require
JWT_SECRET=your-secret-key
GIN_MODE=debug
```

3. Install dependencies

```bash
make init
```

## Running

With Make:

```bash
make build   # compile
make run     # run binary
```

Without Make:

```bash
go run ./cmd/api/
```

With Docker:

```bash
docker build -t restapis-go .
docker run -p 8080:8080 --env-file .env restapis-go
```

## Endpoints

| Method | Route | Auth | Description |
|--------|-------|------|-------------|
| GET | `/healthy` | No | Health check |
| POST | `/api/v1/login` | No | Login and receive JWT token |
| GET | `/api/v1/` | Yes | List all users |
| GET | `/api/v1/:ID` | Yes | Get user by ID |
| POST | `/api/v1/post` | Yes | Create user |
| PATCH | `/api/v1/update/:ID` | Yes | Update user name |
| DELETE | `/api/v1/delete/:ID` | Yes | Delete user |

Protected routes require the `Authorization` header:

```
Authorization: Bearer <token>
```

## Swagger

After running the project, access:

```
http://localhost:8080/swagger/index.html
```

To regenerate docs:

```bash
make docs
```

## Project Structure

```
├── cmd/
│   └── api/
│       ├── main.go
│       ├── routes.go
│       └── controllers/
│           ├── auth.go
│           └── user-controller.go
├── docs/
├── internal/
│   ├── entities/
│   │   └── user.go
│   ├── error/
│   │   └── error.go
│   ├── middleware/
│   │   ├── auth.go
│   │   └── logger.go
│   ├── repository/
│   │   └── repository.go
│   └── use-cases/
│       ├── auth-user.go
│       ├── create-user.go
│       ├── delete-user.go
│       ├── get-user.go
│       └── update-user.go
├── pkg/
│   └── validateName.go
├── .env
├── Dockerfile
├── go.mod
└── MakeFile
```

## Tests

Tests use SQLite in-memory and do not require a running database.

```bash
go get gorm.io/driver/sqlite
go test ./... -v
```

To check coverage:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```