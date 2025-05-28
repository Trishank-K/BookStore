# Book Management System

A Go application for managing a bookstore inventory using MySQL database.

## Setup Instructions

### Development Environment

1. Clone the repository
2. Copy `.env.example` to `.env` and update with your database credentials:
   ```bash
   cp .env.example .env
   ```
3. Edit the `.env` file with your database configuration
4. Run the application:
   ```bash
   go run cmd/main/main.go
   ```

### Production Environment

For production deployment, you need to set the following environment variables:

```bash
# Database Configuration
export DB_USERNAME=your_production_db_username
export DB_PASSWORD=your_production_db_password
export DB_HOST=your_production_db_host
export DB_PORT=your_production_db_port
export DB_DATABASE=your_production_db_name

# Application Configuration
export PORT=8000
export GO_ENV=production
```

You can set these environment variables in your deployment platform (Docker, Kubernetes, Heroku, etc.)

## API Endpoints

- `GET /api/books` - List all books
- `GET /api/books/{id}` - Get a book by ID
- `POST /api/books` - Create a new book
- `PUT /api/books/{id}` - Update a book
- `DELETE /api/books/{id}` - Delete a book

## Dependencies

- GORM - ORM library for Go
- Gorilla Mux - HTTP router and URL matcher
- GoDotEnv - .env file parser
