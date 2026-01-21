# Стадия сборки
FROM golang:latest AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы модулей сначала (для лучшего кэширования)
COPY go.mod go.sum Makefile ./
RUN make deps

# Копируем остальные файлы проекта
COPY . .

# Компилируем приложение с оптимизациями
RUN make build

# Стадия запуска
FROM debian:bookworm-slim AS runtime

# Устанавливаем рабочую директорию
WORKDIR /root/

# Устанавливаем зависимости времени выполнения
RUN apk --no-cache add ca-certificates

# Копируем скомпилированное приложение из стадии builder
COPY --from=builder /app/kafka-consumer-producer .

# Проверяем бинарник
RUN ldd ./kafka-consumer-producer || true 

# Создаем непривилегированного пользователя для безопасности
RUN adduser -D -s /bin/sh appuser
USER appuser

# Запускаем приложение
CMD ["./kafka-consumer-producer"]