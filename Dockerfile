# Этап сборки и тестирования
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Установка необходимых инструментов для тестирования и сборки
RUN apk add --no-cache gcc musl-dev git

# Инициализация модуля и настройка GOPROXY
ENV GOPROXY=direct
ENV GO111MODULE=on

# Копируем только go.mod сначала
COPY go.mod ./

# Скачиваем зависимости и создаем go.sum
RUN go mod download && go mod verify

# Копируем исходный код
COPY . .

# Запускаем тесты
RUN go test -v ./...

# Собираем приложение только если тесты прошли успешно
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Финальный этап
FROM alpine:latest

WORKDIR /app

# Копируем бинарный файл из этапа сборки
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/.env .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]