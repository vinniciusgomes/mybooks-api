# Stage 1: Build stage
FROM golang:1.22.3-alpine AS build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/api

# Stage 2: Final stage
FROM scratch

# Copy the binary from the build stage
COPY --from=build /app/myapp /app/myapp

# Set the entrypoint command
ENTRYPOINT ["/app/myapp"]
