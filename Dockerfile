# Этап сборки и тестирования
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Установка необходимых инструментов для тестирования и сборки
RUN apk add --no-cache gcc musl-dev git make

# Инициализация модуля и настройка GOPROXY
ENV GOPROXY=direct
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

# Копируем только go.mod сначала
COPY go.mod ./

# Скачиваем зависимости и создаем go.sum
RUN go mod download && \
    go get github.com/joho/godotenv && \
    go get github.com/lib/pq && \
    go mod tidy

# Копируем исходный код
COPY . .

# Запускаем тесты
RUN go test -v ./internal/tests/...

# Собираем приложение
RUN go build -o main .

# Финальный этап
FROM alpine:latest

WORKDIR /app

# Копируем бинарный файл и статические файлы из этапа сборки
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/.env .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]