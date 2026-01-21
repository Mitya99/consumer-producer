# Стадия сборки
FROM golang:latest AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /workspace

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
WORKDIR /

# Устанавливаем зависимости времени выполнения
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Копируем скомпилированное приложение из стадии builder
COPY --from=builder /workspace/kafka-consumer-producer /kafka-consumer-producer

# Проверяем бинарник
RUN ldd ./kafka-consumer-producer || true 

# Создаем непривилегированного пользователя для безопасности
RUN groupadd -r appuser && useradd -r -g appuser appuser
USER appuser

ENTRYPOINT ["/kafka-consumer-producer"]