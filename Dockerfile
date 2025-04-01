FROM golang:1.20-alpine AS builder

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api

# Imagen final
FROM alpine:latest

# Instalar dependencias
RUN apk --no-cache add ca-certificates tzdata

# Establecer zona horaria
ENV TZ=America/Bogota

# Crear un usuario no privilegiado
RUN adduser -D -h /app appuser

# Establecer directorio de trabajo
WORKDIR /app

# Copiar el ejecutable desde la etapa de compilación
COPY --from=builder /app/main .

# Cambiar propietario del ejecutable
RUN chown -R appuser:appuser /app

# Cambiar al usuario no privilegiado
USER appuser

# Exponer puerto
EXPOSE 8080

# Comando de inicio
CMD ["./main"]

