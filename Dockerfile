# Etapa 1: Build
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV CGO_ENABLED=1
RUN go build -o backend_golang main.go

# Etapa 2: Imagen final también en Alpine
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/backend_golang .

COPY users.db ./

EXPOSE 8080

CMD ["./backend_golang"]
