# Use Go 1.24.3 as the base image
FROM golang:1.24.3

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Expose port 8000 (adjust this if your server uses another port)
EXPOSE 8000

# Run the application using go run
CMD ["go", "run", "cmd/main.go"]
