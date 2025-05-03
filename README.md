# Grawler

A web crawler service built with Go, Gin, and PostgreSQL. This service allows you to manage URLs and create crawling tasks for them.

## Features

- URL Management
  - Create and store URLs
  - Fetch all stored URLs
  - Configure crawling depth for each URL
- Task Management
  - Create crawling tasks for URLs
  - Track task status (pending, running, completed, failed)
- RESTful API
  - Built with Gin framework
  - Swagger documentation
- Database
  - PostgreSQL for data persistence
  - GORM for database operations

## Tech Stack

- **Backend**: Go 1.24.1
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Containerization**: Docker & Docker Compose
- **Development**: Air (hot reload)

## Project Structure

```
.
├── handlers/     # HTTP request handlers
├── middleware/   # Gin middleware
├── models/       # Database models
├── repository/   # Database operations
├── router/       # API routes
├── Dockerfile    # Backend container configuration
└── docker-compose.yaml  # Multi-container setup
```

## Models

### URL
- `ID`: Unique identifier
- `Url`: The URL to crawl
- `MaxDepth`: Maximum crawling depth
- `IsDisabled`: URL status
- `CreatedAt`: Creation timestamp
- `UpdatedAt`: Last update timestamp

### Task
- `ID`: Unique identifier
- `Status`: Task status (pending/running/completed/failed)
- `UrlID`: Reference to the URL
- `CreatedAt`: Creation timestamp
- `UpdatedAt`: Last update timestamp

## API Endpoints

- `GET /healthz`: Health check
- `POST /url`: Create a new URL
- `GET /urls`: Fetch all URLs
- `POST /crawl`: Create crawling tasks for all URLs

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go 1.24.1 (for local development)

### Running with Docker

1. Clone the repository
2. Build and start the containers:
   ```bash
   docker-compose up --build
   ```
3. The API will be available at `http://localhost:8080`

### Environment Variables

The following environment variables are required:

- `DB_HOST`: Database host (default: db)
- `DB_PORT`: Database port (default: 5432)
- `DB_USER`: Database user (default: grawler)
- `DB_PASSWORD`: Database password (default: grawlerpass)
- `DB_NAME`: Database name (default: grawlerdb)

## Development

For local development:

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Start the development server with hot reload:
   ```bash
   go run github.com/air-verse/air
   ```

## License

MIT
