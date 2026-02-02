# 🚀 Fiber Boilerplate

> Production-ready Go backend boilerplate built with Fiber framework, featuring MongoDB integration, JWT authentication, Redis caching, and automated Swagger documentation.

[![Go Version](https://img.shields.io/badge/Go-1.25.6-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Fiber Version](https://img.shields.io/badge/Fiber-v2.52.11-00ACD7?style=flat)](https://gofiber.io/)
[![MongoDB](https://img.shields.io/badge/MongoDB-Latest-47A248?style=flat&logo=mongodb)](https://www.mongodb.com/)
[![Redis](https://img.shields.io/badge/Redis-Latest-DC382D?style=flat&logo=redis)](https://redis.io/)
[![Swagger](https://img.shields.io/badge/Swagger-Enabled-85EA2D?style=flat&logo=swagger)](https://swagger.io/)

---

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Getting Started](#-getting-started)
- [Configuration](#-configuration)
- [API Documentation](#-api-documentation)
- [Authentication](#-authentication)
- [Scripts](#-scripts)
- [Development](#-development)

---

## ✨ Features

- ⚡ **High Performance** - Built on Fiber (Express-inspired framework for Go)
- 🔐 **JWT Authentication** - Secure Bearer & Basic authentication
- 📊 **MongoDB Integration** - NoSQL database with connection pooling
- 🚀 **Redis Caching** - High-performance data caching
- 📝 **Auto Swagger Docs** - Interactive API documentation
- 🏗️ **Clean Architecture** - Modular structure with separation of concerns
- 🔄 **Auto Port Management** - Intelligent startup script that handles port conflicts
- 🛡️ **Error Handling** - Centralized error handling middleware
- 📦 **Environment Config** - Secure environment variable management
- 🎯 **Query Builder** - MongoDB query abstraction layer

---

## 🛠️ Tech Stack

### Core Framework
- **[Fiber v2](https://gofiber.io/)** - Express-inspired web framework written in Go
- **Go 1.25.6** - Modern, efficient programming language

### Database & Caching
- **[MongoDB](https://www.mongodb.com/)** - NoSQL database for flexible data storage
- **[Redis](https://redis.io/)** - In-memory data structure store for caching

### Authentication & Security
- **[JWT](https://jwt.io/)** - JSON Web Tokens for secure authentication
- **golang-jwt/jwt/v5** - JWT implementation for Go

### Documentation
- **[Swagger/OpenAPI](https://swagger.io/)** - Automated API documentation
- **swaggo/swag** - Swagger documentation generator
- **gofiber/swagger** - Fiber Swagger middleware

### Development Tools
- **godotenv** - Environment variable management
- **Custom Scripts** - Automated startup and development workflows

---

## 📁 Project Structure

```
fiber-_boilerplate/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── app/
│   │   ├── app.go             # App initialization
│   │   ├── middleware.go      # Middleware registration
│   │   └── routes.go          # Route registration
│   ├── config/
│   │   └── config.go          # Environment configuration
│   ├── error/
│   │   └── error.go           # Error handling
│   ├── lib/
│   │   ├── dbConnection/      # MongoDB connection
│   │   ├── redis/             # Redis client
│   │   └── swagger/           # Swagger setup
│   ├── middleware/
│   │   ├── basicAuth.go       # Basic authentication
│   │   └── bearerAuth.go      # JWT Bearer authentication
│   ├── modules/
│   │   └── user/              # User module
│   │       ├── v1/
│   │       │   ├── userController.go  # HTTP handlers
│   │       │   ├── userService.go     # Business logic
│   │       │   ├── userRepository.go  # Data access
│   │       │   ├── userDto.go         # Data transfer objects
│   │       │   └── userRoute.go       # Route definitions
│   │       ├── userModel.go           # User entity
│   │       └── index.go
│   ├── querybuilder/          # MongoDB query builder
│   └── utils/                 # Utility functions
├── scripts/
│   └── start.sh               # Smart startup script
├── docs/                      # Auto-generated Swagger docs
├── .env                       # Environment variables
├── go.mod                     # Go dependencies
└── README.md                  # This file
```

---

## 🚀 Getting Started

### Prerequisites

- **Go** 1.25.6 or higher
- **MongoDB** running on `localhost:27017`
- **Redis** (optional, for caching features)
- **Swagger CLI** for documentation generation

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/addixit1/fiber-boilerplate.git
   cd fiber-_boilerplate
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Install Swagger CLI**
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

4. **Configure environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

5. **Make start script executable**
   ```bash
   chmod +x scripts/start.sh
   ```

### Running the Application

**Option 1: Using the Smart Start Script (Recommended)**
```bash
./scripts/start.sh
```
This script automatically:
- Kills any process running on port 3010
- Generates Swagger documentation
- Starts the server

**Option 2: Manual Steps**
```bash
# Generate Swagger docs
swag init -g cmd/main.go
# OR
$(go env GOPATH)/bin/swag init -g cmd/main.go

# Run the server
go run cmd/main.go
```

### Verify Installation

- **Server**: http://localhost:3010
- **Swagger UI**: http://localhost:3010/swagger/index.html
- **Health Check**: `GET http://localhost:3010/api/v1/health` (if implemented)

---

## ⚙️ Configuration

### Environment Variables

Create a `.env` file in the root directory:

```env
ENV=development
PORT=3010
MONGO_URI=mongodb://localhost:27017
MONGO_DB=fiber_db
JWT_SECRET=your_super_secret_key_here
REDIS_HOST=localhost:6379
```

### Configuration Details

| Variable | Description | Default |
|----------|-------------|---------|
| `ENV` | Environment (development/production) | `development` |
| `PORT` | Server port | `3010` |
| `MONGO_URI` | MongoDB connection string | `mongodb://localhost:27017` |
| `JWT_SECRET` | Secret key for JWT signing | `supersecret` |
| `REDIS_HOST` | Redis server address | `localhost:6379` |

---

## 📝 API Documentation

### Swagger/OpenAPI

Interactive API documentation is automatically generated and available at:

**http://localhost:3010/swagger/index.html**

### Swagger Annotations

The project uses Swagger annotations in code comments:

```go
// @title Fiber Production API
// @version 1.0
// @description Production ready Fiber backend
// @host localhost:3010
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

### Regenerating Documentation

```bash
swag init -g cmd/main.go
```

---

## 🔐 Authentication

### Authentication Methods

1. **Bearer Token (JWT)**
   ```bash
   curl -H "Authorization: Bearer YOUR_JWT_TOKEN" \
        http://localhost:3010/api/v1/protected
   ```

2. **Basic Auth**
   ```bash
   curl -u username:password \
        http://localhost:3010/api/v1/basic-protected
   ```

### JWT Token Structure

```json
{
  "user_id": "string",
  "email": "string",
  "exp": 1234567890
}
```

---

## 📜 Scripts

### `scripts/start.sh`

Smart startup script with the following features:

- **Port Conflict Resolution**: Automatically detects and kills processes on port 3010
- **Doc Generation**: Regenerates Swagger documentation
- **Error Handling**: Validates each step before proceeding
- **Color-Coded Output**: Clear visual feedback

**Usage:**
```bash
./scripts/start.sh
```

**What it does:**
```
🔍 Checking if port 3010 is already in use...
✅ Port 3010 is free!

📝 Generating Swagger documentation...
✅ Swagger docs generated successfully!

🚀 Starting Fiber server...
📍 Swagger UI: http://localhost:3010/swagger/index.html
```

---

## 🔧 Development

### Project Architecture

The project follows **Clean Architecture** principles:

```
Controller → Service → Repository → Database
     ↓          ↓           ↓
    DTO    Business      Query
           Logic         Builder
```

### Adding a New Module

1. Create module directory: `internal/modules/yourmodule/`
2. Create version subdirectory: `internal/modules/yourmodule/v1/`
3. Implement the following files:
   - `yourmoduleModel.go` - Entity definition
   - `yourmoduleController.go` - HTTP handlers
   - `yourmoduleService.go` - Business logic
   - `yourmoduleRepository.go` - Data access
   - `yourmoduleDto.go` - DTOs
   - `yourmoduleRoute.go` - Route definitions
4. Register routes in `internal/app/routes.go`

### Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use meaningful variable names
- Add Swagger comments for all endpoints
- Write unit tests for services and repositories

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/modules/user/v1/...
```

---

## 🗂️ Key Dependencies

```go
github.com/gofiber/fiber/v2       // Web framework
github.com/gofiber/swagger        // Swagger middleware
github.com/golang-jwt/jwt/v5      // JWT authentication
github.com/joho/godotenv          // Environment config
github.com/redis/go-redis/v9      // Redis client
github.com/swaggo/swag            // Swagger generator
go.mongodb.org/mongo-driver       // MongoDB driver
```

---

## 🐛 Troubleshooting

### Port Already in Use

The startup script handles this automatically. If running manually:

```bash
# Find process on port 3010
lsof -ti:3010

# Kill the process
kill -9 $(lsof -ti:3010)
```

### Swagger Docs Not Generating

```bash
# Check if swag is installed
which swag
# OR
$(go env GOPATH)/bin/swag --version

# Add GOPATH bin to PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

### MongoDB Connection Issues

- Ensure MongoDB is running: `sudo systemctl status mongod`
- Check connection string in `.env`
- Verify MongoDB port: `27017`

---

## 📄 License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Author

**Aman Dixit**
- Email: aman.dixitiimt@gmail.com
- GitHub: [@addixit1](https://github.com/addixit1)

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 🙏 Acknowledgments

- [Fiber Framework](https://gofiber.io/)
- [MongoDB](https://www.mongodb.com/)
- [Swagger](https://swagger.io/)
- Go Community

---

<div align="center">

**⭐ Star this repository if you find it helpful!**

Made with ❤️ using Go and Fiber

</div>
