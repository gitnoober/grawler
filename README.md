# Grawler

**Grawler** is an open-source web crawler and content summarizer powered by Go, Gin, PostgreSQL, and local LLMs via Ollama.
Easily manage URLs, automate crawling tasks, and generate human-quality summaries of web content using state-of-the-art language models—all in your own infrastructure.
🚀 Fast, concurrent crawling with task management
🤖 LLM-powered content summaries (Ollama integration)
🗂️ RESTful API for easy integration
🐳 Fully containerized with Docker & Compose
Perfect for research, monitoring, or building your own knowledge base from the web!

---

## 🚀 Features

- **URL Management**: Create, list, configure depth, enable/disable URLs.
- **Task Management**: Generate crawling tasks, track statuses (*pending*, *running*, *completed*, *failed*).
- **Crawling Engine**: Concurrent workers fetch page content with timeout and retry logic.
- **Summarization**: Background summarizer calls an LLM API to generate concise content summaries.
- **LLM-Powered Summaries**: Uses a local LLM (via Ollama) to generate human-quality summaries of crawled content.
- **RESTful API**: JSON endpoints powered by Gin, with built-in Swagger documentation.
- **Containerized**: Docker & Docker Compose for easy setup and scaling.
- **Hot Reload**: Development with Air for instant reloads on code changes.

---

## 🔧 Tech Stack

| Component      | Technology         |
| -------------- | ------------------ |
| Language       | Go 1.24.1          |
| Web Framework  | Gin                |
| ORM            | GORM (PostgreSQL)  |
| Container      | Docker & Docker Compose |
| Hot Reload     | Air                |

---

## 📁 Project Structure

```
.
├── handlers/          # HTTP request handlers
├── middleware/        # Gin middleware for dependency injection
├── models/            # GORM models (URL, Task, TaskResponse)
├── repository/        # Database repositories
├── queue/             # Task queue and crawling workers
├── summarizer/        # Background summarization worker
├── router/            # API route definitions
├── main.go            # Application entrypoint
├── Dockerfile         # Backend Docker configuration
└── docker-compose.yaml # Multi-container setup
```

---

## ⚙️ Prerequisites

- Go 1.24+ installed (for local dev)
- Docker & Docker Compose (for containerized setup)
- PostgreSQL database

---

## 🏁 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/gitnoober/grawler.git
cd grawler
```

### 2. Set Environment Variables

Create a `.env` file or export in your shell:

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=grawler
export DB_PASSWORD=grawlerpass
export DB_NAME=grawlerdb
export PORT=8080
export LLM_API_URL=http://localhost:11434/api/generate
```

#### Ollama LLM Setup

Ensure you have the Ollama CLI installed for local LLM inference:

```bash
# macOS
brew install ollama
# or visit https://ollama.ai for other platforms

# Pull a model (e.g., llama2)
ollama pull llama2:latest

# Start the Ollama server
ollama serve llama2 --port 11434
```

### 3. Run with Docker

```bash
docker-compose up --build
```

- **Backend**: http://localhost:8080
- **PostgreSQL**: localhost:5432

### 4. Local Development

```bash
# Install dependencies
go mod download

# Start server with hot reload
air
```

---

## 📡 API Endpoints

| Method | Path               | Description                          |
| ------ | ------------------ | ------------------------------------ |
| GET    | `/healthz`         | Health check endpoint                |
| POST   | `/url`             | Create a new URL to crawl            |
| GET    | `/urls`            | Fetch all stored URLs                |
| POST   | `/crawl`           | Generate crawling tasks for all URLs |
| GET    | `/summaries`       | List all URL summaries               |
| GET    | `/summaries/:id`   | Fetch a summary by its ID            |

---

## 📄 License

This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for details.
