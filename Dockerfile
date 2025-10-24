# ====== STAGE 1: Build binary ======
FROM golang:1.23.2 AS builder

WORKDIR /app

# Copy dependency files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o app .

# ====== STAGE 2: Run binary on minimal image ======
FROM debian:bookworm-slim

WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /app/app .

COPY --from=builder /app/assets_static ./assets_static

# Install CA certs for HTTPS
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*

# Create folder for uploads
RUN mkdir -p /app/assets


CMD ["./app"]
