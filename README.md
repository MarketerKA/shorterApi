# URL Shortener Service

Сервис для создания коротких URL-ссылок с использованием Go и PostgreSQL/In-Memory хранилища.

## Возможности

- Создание коротких ссылок
- Получение оригинальных ссылок по короткому коду
- Веб-интерфейс для удобного использования
- Поддержка разных типов хранилищ (PostgreSQL/In-Memory)
- Валидация URL
- История запросов
- Обработка ошибок

## Структура проекта 

```
.
├── main.go                 # Точка входа в программу
├── static/                 # Статические файлы
│   └── test.html          # Веб-интерфейс
└── internal/              # Внутренние пакеты
    ├── handler.go         # HTTP обработчики
    ├── service.go         # Бизнес-логика
    ├── validator/         # Валидация данных
    ├── errors/           # Обработка ошибок
    └── storage/          # Хранилище данных
        ├── storage.go    # Интерфейс хранилища
        ├── postgres.go   # PostgreSQL реализация
        ├── memory.go     # In-Memory реализация
        └── factory.go    # Фабрика хранилищ
```

## Установка и запуск

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```

2. Создайте файл .env с настройками:
   ```env
   STORAGE_TYPE=memory    # или postgres
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=urlshortener
   DB_SSLMODE=disable
   ```

3. Если используете PostgreSQL, создайте базу данных:
   ```sql
   CREATE DATABASE urlshortener;
   ```

4. Установите зависимости:
   ```bash
   go mod tidy
   ```

5. Запустите сервер:
   ```bash
   go run main.go
   ```

6. Откройте веб-интерфейс:
   ```
   http://localhost:8080/static/test.html
   ```

## API Endpoints

### Создание короткой ссылки
```bash
POST /create
Content-Type: application/json

{
    "original_url": "https://example.com"
}
```

### Получение информации о ссылке
```bash
GET /info/{short_code}
```

### Переход по короткой ссылке
```bash
GET /{short_code}
```

## Веб-интерфейс

Веб-интерфейс предоставляет следующие возможности:
- Создание коротких ссылок
- Получение информации о ссылках
- История запросов
- Визуальное отображение ошибок
- Логирование запросов в консоли браузера

## Разработка

При разработке использованы следующие принципы и технологии:
- Clean Architecture
- Dependency Injection
- Interface Segregation
- Factory Pattern
- CORS поддержка
- Обработка ошибок
- Логирование









