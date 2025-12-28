# Стадия сборки
FROM golang:latest AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы модулей сначала (для лучшего кэширования)
COPY go.mod  ./
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Компилируем приложение с оптимизациями
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# Стадия запуска
FROM alpine:latest AS runtime

# Устанавливаем рабочую директорию
WORKDIR /root/

# Устанавливаем зависимости времени выполнения
RUN apk --no-cache add ca-certificates

# Копируем скомпилированное приложение из стадии builder
COPY --from=builder /app/main .

# Создаем непривилегированного пользователя для безопасности
RUN adduser -D -s /bin/sh appuser
USER appuser

# Запускаем приложение
CMD ["./main"]