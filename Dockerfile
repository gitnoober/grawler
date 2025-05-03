FROM golang:latest

WORKDIR /app

# Install air (latest version as of now)
RUN go install github.com/air-verse/air@latest

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Expose whatever port your app runs on
EXPOSE 8080

# Start Air (hot reload dev server)
CMD ["air"]