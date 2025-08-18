# Gunakan image Go untuk build
FROM golang:1.23.2 AS builder

WORKDIR /app

# Copy semua file go mod dan sum dulu (supaya cache build cepat)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o app .

# Image ringan untuk run
FROM debian:bookworm-slim

WORKDIR /root/
COPY --from=builder /app/app .

# Install CA certificates biar TLS jalan
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*
# COPY .env .env

CMD ["./app"]
