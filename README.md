# URL Shortener Service

Сервис для создания коротких URL-ссылок с использованием Go и PostgreSQL.

## Структура проекта 

- `main.go` - точка входа в программу
- `internal/` - папка с внутренними компонентами сервиса
  - `storage.go` - хранилище данных
  - `service.go` - сервис для создания и получения URL-ссылок
  - `handler.go` - обработчик HTTP-запросов

## Установка и запуск

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/MarketerKA/shorterApi.git
   ```

2. Установите зависимости:
   ```bash
   go mod tidy
   ```

3. Запустите сервер:
   ```bash
   go run main.go
   ```

4. Откройте браузер и перейдите по адресу:
   ```bash
   http://localhost:8080
   ```

5. Используйте API для создания и получения URL-ссылок:
   - Для создания короткой ссылки отправьте POST-запрос на `/create` с JSON-телом:
     ```bash
     curl -X POST -H "Content-Type: application/json" -d '{"original_url": "https://example.com"}' http://localhost:8080/create
     ```

6. Получите короткую ссылку, перейдя по адресу:
   ```bash
   http://localhost:8080/
   ```









