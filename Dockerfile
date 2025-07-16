# Etapa 1: Build
FROM golang:1.24 AS builder

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de módulos y descarga dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código fuente
COPY . .

# Compila el binario
RUN CGO_ENABLED=0 GOOS=linux go build -o backend_golang main.go

# Etapa 2: Imagen final ligera
FROM alpine:latest

# Instalar certificado y sqlite runtime (si necesitas sqlite)
RUN apk add --no-cache ca-certificates sqlite


WORKDIR /root/

# Copia el binario compilado desde la etapa builder
COPY --from=builder /app/backend_golang .

# Copia la base de datos SQLite si la usas
COPY users.db ./

# Expone el puerto que usas en Gin (cambia si es otro)
EXPOSE 8080

# Ejecuta la API al iniciar el contenedor
CMD ["./backend_golang"]

