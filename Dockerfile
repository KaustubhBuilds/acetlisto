# Build stage
FROM golang:1.27-trixie AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application (IMPORTANT FIX HERE)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Final stage
FROM gcr.io/distroless/static-debian11

ENV PORT=8080
EXPOSE 8080

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Run as non-root user
USER nonroot:nonroot

# Command to run
CMD ["./main"]
