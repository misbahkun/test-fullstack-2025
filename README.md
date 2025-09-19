# Test Fullstack 2025

This repository contains two separate Go applications as part of a fullstack test:

1. **tugas1-faktorial** - A simple Go application that calculates a mathematical function
2. **tugas2-fiber-login** - A REST API for user authentication using Fiber framework and Redis

## Project Structure

```
.
├── tugas1-faktorial/
│   ├── go.mod
│   └── main.go
└── tugas2-fiber-login/
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── config/
    │   └── database.go
    ├── handler/
    │   └── auth_handler.go
    └── model/
        └── auth_model.go
```

## Tugas 1: Faktorial

A Go application that calculates a mathematical function F(n) where:
- F(n) = ceil(n! / 2^n)

### How to Run

1. Navigate to the `tugas1-faktorial` directory:
   ```bash
   cd tugas1-faktorial
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

### Function Details

The application contains two functions:
1. `calculateFactorial(n int) float64` - Calculates the factorial of n
2. `F(n int) float64` - Calculates the complete function F(n) = ceil(n! / 2^n)

By default, the program calculates F(10) and F(5) and prints the results.

## Tugas 2: Fiber Login API

A REST API for user authentication built with the Fiber framework and Redis as the database.

### Features

- User login endpoint
- SHA1 password hashing
- Redis as storage
- Docker Compose setup for Redis

### Prerequisites

- Go 1.25+
- Docker and Docker Compose (for Redis)

### How to Run

1. Navigate to the `tugas2-fiber-login` directory:
   ```bash
   cd tugas2-fiber-login
   ```

2. Start Redis using Docker Compose:
   ```bash   
   docker-compose up -d
   ```

3. Download Go dependencies:
   ```bash
   go mod tidy
   ```

4. Run the Go application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:3000`

### API Endpoints

#### POST `/api/v1/login`

Login endpoint that accepts a JSON body with username and password.

**Request Body:**
```json
{
  "username": "doni",
  "password": "password123"
}
```

**Response (Success):**
```json
{
  "status": "success",
  "message": "Login berhasil!",
  "user": {
    "realname": "Aberto Doni Sianturi",
    "email": "adss@gmail.com"
  }
}
```

**Response (Error):**
```json
{
  "status": "error",
  "message": "Username atau password salah."
}
```

### Default User

The application seeds the database with a default user:
- Username: `doni`
- Password: `password123`
- Real Name: `Aberto Doni Sianturi`
- Email: `adss@gmail.com`

### Project Structure

```
tugas2-fiber-login/
├── main.go              # Application entry point
├── config/
│   └── database.go      # Redis connection and seeding
├── handler/
│   └── auth_handler.go  # Authentication endpoint handler
└── model/
    └── auth_model.go    # Data structures for authentication
```

## Dependencies

### Tugas 1
- Standard Go library (math)

### Tugas 2
- github.com/gofiber/fiber/v2 - Web framework
- github.com/redis/go-redis/v9 - Redis client
- Standard Go library (crypto/sha1, encoding/json, context, etc.)

## License

This project is for educational/testing purposes only.