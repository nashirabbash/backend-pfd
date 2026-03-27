# PFD Backend - Go REST API with WebSocket

A complete backend API implementation using Go Fiber, PostgreSQL, and WebSocket support following Clean Architecture principles.

## 🚀 Project Overview

This project implements a complete backend solution with 8 implementation phases:
- Phase 1: Go module initialization and dependency management
- Phase 2: Database configuration and connection setup
- Phase 3: Model structure and auto-migration
- Phase 4: Request validation system
- Phase 5: JWT authentication and authorization
- Phase 6: REST API endpoints
- Phase 7: WebSocket integration
- Phase 8: Application entry point and assembly

## 🛠️ Tech Stack

| Technology | Package | Version |
|-----------|---------|---------|
| Language | Go | v1.25.0+ |
| Web Framework | Fiber | v2.52.12 |
| Database | PostgreSQL | - |
| ORM | GORM | v1.31.1 |
| WebSocket | Fiber WebSocket | v1.3.4 |
| Validator | Go Validator | v10.30.1 |
| Authentication | JWT | v5.3.1 |
| Environment | godotenv | v1.5.1 |

## 📂 Project Structure

```
pfd-be/
├── cmd/
│   └── api/
│       └── main.go                 # Entry point
├── internal/
│   ├── config/
│   │   └── config.go               # Configuration management
│   ├── database/
│   │   ├── database.go             # Database connection
│   │   └── migration.go            # Auto-migration
│   ├── model/
│   │   └── user.go                 # User model
│   ├── dto/
│   │   └── auth.go                 # Data Transfer Objects
│   ├── repository/
│   │   └── user.go                 # Data access layer
│   ├── service/
│   │   └── auth.go                 # Business logic
│   ├── handler/
│   │   └── handlers.go             # HTTP handlers
│   ├── middleware/
│   │   ├── jwt.go                  # JWT utilities
│   │   ├── auth.go                 # Auth middleware
│   │   └── validator.go            # Validation middleware
│   └── route/
│       └── routes.go               # Route definitions
├── .env.example                    # Environment variables template
├── go.mod                          # Module definition
├── go.sum                          # Dependency checksums
└── README.md                       # This file
```

## 🔧 Setup Instructions

### Prerequisites
- Go 1.24+
- PostgreSQL 12+
- Git

### Installation

1. **Clone the repository**
   ```bash
  git clone <your-repo-url>
  cd pfd-be
   ```

2. **Configure environment variables**
  Copy `.env.example` to `.env`, then adjust values:
   ```env
  DATABASE_URL=postgres://postgres:postgres@localhost:5432/pfd_db?sslmode=disable
   PORT=3000
   ENV=development
  JWT_SECRET=change-this-secret
   JWT_EXPIRATION=24
   ```

3. **Create PostgreSQL database**
   ```bash
   createdb pfd_db
   ```

4. **Run the application**
   ```bash
   go run ./cmd/api/main.go
   ```
   Or use the pre-built binary:
   ```bash
   ./bin/api
   ```

## 📡 API Endpoints

### Authentication Endpoints

#### Register User
```
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe"
}

Response: 201 Created
{
  "token": "eyJhbGc...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

#### Login User
```
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}

Response: 200 OK
{
  "token": "eyJhbGc...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

#### Get Current User (Protected)
```
GET /api/auth/me
Authorization: Bearer <your_token_here>

Response: 200 OK
{
  "user_id": 1,
  "email": "user@example.com"
}
```

### WebSocket Endpoint
```
WS /ws

Connection: websocket protocol
Message Format: plain text (echo test)
Response: Echo: <your_message>
```

### Health Check
```
GET /health

Response: 200 OK
{
  "status": "OK"
}
```

## 🔐 Authentication

The API uses JWT (JSON Web Tokens) for authentication:
- Token generated during registration and login
- Set token in Authorization header for protected routes:
  ```
  Authorization: Bearer <your_token_here>
  ```
- Token expiration: Default 24 hours (configurable via `JWT_EXPIRATION`)

## 🧪 Testing the API

### Using cURL

**Register:**
```bash
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123",
    "name": "Test User"
  }'
```

**Login:**
```bash
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

**Health Check:**
```bash
curl http://localhost:3000/health
```

**WebSocket (using wscat):**
```bash
npm install -g wscat
wscat -c ws://localhost:3000/ws
```

## ✅ Definition of Done - Verification Checklist

- [x] Code compiles and runs with `go run ./cmd/api/main.go`
- [x] Database PostgreSQL connection successful
- [x] `POST /api/auth/register` returns JWT token
- [x] `POST /api/auth/login` returns JWT token
- [x] Protected route `GET /api/auth/me` rejects requests without valid token (401)
- [x] WebSocket endpoint (`/ws`) establishes connection
- [x] WebSocket sends echo response to messages
- [x] Health check endpoint works (`/health`)

## 🏗️ Architecture

The project follows **Clean Architecture** patterns:

```
Handler (HTTP)
    ↓
Middleware (Validation, Auth)
    ↓
Service (Business Logic)
    ↓
Repository (Data Access)
    ↓
Database (PostgreSQL + GORM)
```

## 📝 Key Features

✅ **JWT Authentication** - Secure token-based authentication  
✅ **Request Validation** - Built-in validator for all DTOs  
✅ **WebSocket Support** - Real-time bidirectional communication  
✅ **Database Migration** - Automatic schema creation with GORM auto-migrate  
✅ **Error Handling** - Comprehensive error responses  
✅ **CORS Support** - Cross-origin resource sharing enabled  
✅ **Logging** - Request and response logging middleware  
✅ **Clean Code** - Organized by layers and packages  

## 🚀 Next Steps / Future Enhancements

- Add rate limiting
- Implement refresh token mechanism
- Add database transaction support
- Implement audit logging
- Add file upload functionality
- Create Swagger/OpenAPI documentation
- Add comprehensive test suite
- Implement session management
- Add role-based access control (RBAC)

## 📚 References

- [Go Fiber Documentation](https://docs.gofiber.io/)
- [GORM Documentation](https://gorm.io/)
- [JWT Go Library](https://github.com/golang-jwt/jwt)
- [Go Validator](https://github.com/go-playground/validator)

## 📄 License

This project is part of the PFD Backend initiative.

---

**Status:** ✅ Implementation Complete  
**Last Updated:** March 28, 2026  
**Go Version:** v1.25.0
