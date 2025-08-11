# Go-Project-Template

A Go-based REST API service providing authentication, token management, and database connectivity features.

## Features

- JWT-based authentication and authorization
- PostgreSQL database integration with GORM
- Token creation, validation, and time extension
- Email functionality with SMTP support
- CORS-enabled REST API endpoints
- Multi-schema database support (common, accounts, sitlpos, iservice)
- Configuration-based authorization bypass

## Prerequisites

- Go 1.16 or higher
- PostgreSQL database
- Access to required database schemas

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Shawan-Das/Go-Project-Template.git
cd Go-Project-Template
```

2. Install dependencies:
```bash
go mod tidy
```

3. Configure the application by editing `config/local-config.json` with your database and SMTP settings.

## Configuration

Update `config/local-config.json` with your specific settings:

```json
{
  "db_host": "your-database-host",
  "db_port": 5432,
  "db_name": "your-database-name",
  "db_username": "your-username",
  "db_password": "your-password",
  "access_secret": "your-jwt-secret-key",
  "api_port": "3000",
  "timeout": 60
}
```

## Usage

### Building the Application

```bash
go build
```

### Running the Service

```bash
./iservice <port> config/local-config.json
```

Example:
```bash
./iservice 8080 config/local-config.json
```

### Development Commands

Format code:
```bash
gofumpt -l -w .
```

Build and run:
```bash
gofumpt -l -w . && go build && ./iservice 8080 config/local-config.json
```

### Cross-platform Build (Linux)

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags "-w -extldflags '-static'"
```

## API Endpoints

- `GET /` - Health check
- `POST /ValidateToken` - Token validation
- `POST /ValidateToken_V2` - Enhanced token validation
- Authentication endpoints for login and token management

## Project Structure

```
├── config/           # Configuration files
├── dto/              # Data Transfer Objects
├── model/            # Database models
├── repository/       # Data access layer
├── service/          # Business logic layer
├── util/             # Utility functions
└── main.go           # Application entry point
```

## Database Schemas

The application supports multiple PostgreSQL schemas:
- `common` - Common application data
- `accounts` - Account management

## Authentication

The service uses JWT tokens for authentication with configurable bypass routes for public endpoints. Token expiration and refresh functionality is built-in.

## License

This project is licensed under the terms specified in the license file.
