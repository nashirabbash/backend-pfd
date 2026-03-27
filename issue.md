# Issue: Setup Go Backend Project (PFD-BE)

## Objective

Buat project Go backend dari nol di folder ini. Project ini akan menjadi backend API server dengan fitur real-time menggunakan WebSocket.

---

## Tech Stack

| Kategori      | Library / Tool                                                                                                            |
| ------------- | ------------------------------------------------------------------------------------------------------------------------- |
| Language      | Go (latest stable)                                                                                                        |
| Web Framework | [Go Fiber](https://github.com/gofiber/fiber) v2                                                                           |
| Database      | PostgreSQL                                                                                                                |
| ORM           | [GORM](https://gorm.io/) + PostgreSQL driver                                                                              |
| Validation    | [go-playground/validator](https://github.com/go-playground/validator) v10                                                 |
| WebSocket     | [Fiber WebSocket](https://github.com/gofiber/websocket) (gofiber/contrib/websocket)                                       |
| Auth          | JWT menggunakan [golang-jwt](https://github.com/golang-jwt/jwt) v5                                                        |
| Config        | Environment variables (gunakan [godotenv](https://github.com/joho/godotenv) atau [viper](https://github.com/spf13/viper)) |

---

## Task Breakdown

### 1. Inisialisasi Project

- Jalankan `go mod init` dengan module name yang sesuai
- Buat file `.env.example` berisi template environment variables yang dibutuhkan (DB host, port, user, password, db name, JWT secret, app port)
- Buat file `.gitignore` untuk Go project (include `.env`, binary, tmp, dll)
- Install semua dependency yang disebutkan di tech stack

### 2. Struktur Folder

Buat struktur folder berikut:

```
.
├── cmd/
│   └── server/
│       └── main.go              # Entry point aplikasi
├── internal/
│   ├── config/
│   │   └── config.go            # Load env & app configuration
│   ├── database/
│   │   └── database.go          # Koneksi PostgreSQL via GORM
│   ├── middleware/
│   │   ├── auth.go              # JWT auth middleware
│   │   └── cors.go              # CORS middleware
│   ├── model/
│   │   └── user.go              # GORM model (contoh: User)
│   ├── handler/
│   │   ├── auth_handler.go      # Handler untuk register/login
│   │   ├── user_handler.go      # Handler untuk user CRUD
│   │   └── ws_handler.go        # Handler untuk WebSocket
│   ├── service/
│   │   ├── auth_service.go      # Business logic auth (hash password, generate JWT)
│   │   └── user_service.go      # Business logic user
│   ├── repository/
│   │   └── user_repository.go   # Database query layer
│   ├── dto/
│   │   ├── request.go           # Struct request + validation tags
│   │   └── response.go          # Struct response standar
│   ├── validator/
│   │   └── validator.go         # Setup validator instance & custom validation helper
│   └── router/
│       └── router.go            # Setup semua route & group
├── pkg/
│   └── jwt/
│       └── jwt.go               # Utility generate & parse JWT token
├── .env.example
├── .gitignore
├── go.mod
└── go.sum
```

### 3. Konfigurasi & Database

- Buat config loader yang membaca dari `.env` file
- Buat fungsi koneksi database PostgreSQL menggunakan GORM
- Implementasi auto-migrate untuk model yang didefinisikan
- Pastikan koneksi database bisa retry jika gagal saat startup

### 4. Validator Setup

- Buat instance global validator menggunakan `go-playground/validator`
- Buat helper function untuk parse validation error menjadi response yang user-friendly
- Gunakan struct tags pada DTO untuk validation rules (contoh: `validate:"required,email"`)

### 5. JWT Authentication

- Buat utility function untuk:
  - Generate access token (short-lived, misal 15 menit)
  - Generate refresh token (long-lived, misal 7 hari)
  - Parse & validate token
- Buat auth middleware Fiber yang:
  - Mengambil token dari header `Authorization: Bearer <token>`
  - Validate token dan inject user info ke `fiber.Ctx.Locals`
  - Return 401 jika token invalid/expired

### 7. Auth Endpoints

Buat endpoint:

| Method | Path                 | Deskripsi                | Auth                             |
| ------ | -------------------- | ------------------------ | -------------------------------- |
| POST   | `/api/auth/register` | Register user baru       | No                               |
| POST   | `/api/auth/login`    | Login, return JWT tokens | No                               |
| POST   | `/api/auth/refresh`  | Refresh access token     | No (pakai refresh token di body) |

- Password harus di-hash menggunakan bcrypt sebelum disimpan
- Validasi input menggunakan validator
- Return response format yang konsisten

### 8. User Endpoints (Protected)

Buat endpoint (semua butuh JWT auth):

| Method | Path            | Deskripsi                   |
| ------ | --------------- | --------------------------- |
| GET    | `/api/users/me` | Get current user profile    |
| PUT    | `/api/users/me` | Update current user profile |

### 9. WebSocket

- Setup WebSocket upgrade endpoint di `/ws`
- Implementasi basic WebSocket handler yang bisa:
  - Accept koneksi WebSocket
  - Terima pesan dari client
  - Broadcast pesan ke semua client yang terkoneksi
- Buat simple connection manager (track connected clients)
- WebSocket endpoint boleh di-protect dengan JWT (token dikirim via query param saat upgrade)

### 10. Router & Middleware

- Setup Fiber app dengan config dasar (body limit, dll)
- Pasang CORS middleware
- Buat route group `/api/v1` untuk REST endpoints
- Pisahkan route public (auth) dan protected (butuh JWT middleware)
- Pasang recovery middleware untuk handle panic

### 11. Response Format

Gunakan format response JSON yang konsisten:

```json
// Success
{
  "success": true,
  "message": "descriptive message",
  "data": {}
}

// Error
{
  "success": false,
  "message": "error message",
  "errors": []  // optional, untuk validation errors
}
```

### 12. Entry Point (main.go)

`main.go` harus:

- Load config dari env
- Koneksi ke database
- Run auto-migrate
- Setup Fiber app + middleware + router
- Graceful shutdown handling (listen OS signal)
- Start server di port dari env

---

## Catatan Penting

- **Jangan gunakan framework pattern yang terlalu complex.** Cukup layer sederhana: Handler → Service → Repository.
- **Semua secret/credential harus dari environment variable**, jangan hardcode.
- **Error handling harus konsisten** — selalu return JSON response, jangan panic.
- **Gunakan UUID untuk primary key**, bukan auto-increment integer.
- **Password wajib di-hash** sebelum masuk database.
- **WebSocket cukup basic** — yang penting bisa connect, kirim, dan broadcast pesan. Tidak perlu room/channel system dulu.

---

## Definition of Done

- [ ] `go run cmd/server/main.go` berjalan tanpa error
- [ ] Bisa connect ke PostgreSQL dan auto-migrate model
- [ ] Endpoint register & login berfungsi, return JWT
- [ ] Endpoint protected return 401 tanpa token, dan berfungsi normal dengan token valid
- [ ] WebSocket endpoint bisa menerima koneksi dan broadcast pesan
- [ ] Validation error menghasilkan response yang jelas
- [ ] Semua config dibaca dari `.env`
