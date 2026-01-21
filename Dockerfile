# ETAPA 1: Compilación
FROM golang:1.25-alpine AS builder

# Instalar dependencias para sqlite3 (CGO)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copiar archivos de dependencia primero para cachear capas
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar el binario optimizado (quitando símbolos de debug para reducir peso)
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o conspiracy-app main.go

# ETAPA 2: Producción
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

# Copiar solo el binario y las carpetas necesarias desde el builder
COPY --from=builder /app/conspiracy-app .
COPY --from=builder /app/views ./views
COPY --from=builder /app/static ./static

# Exponer puerto
EXPOSE 1323

# Comando de ejecución
CMD ["./conspiracy-app"]
