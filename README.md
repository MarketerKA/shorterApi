# URL Shortener Service

Сервис для создания коротких URL-ссылок с использованием Go и PostgreSQL/In-Memory хранилища.

## Установка и запуск с нуля

### Предварительные требования
1. Установленный Docker: https://docs.docker.com/get-docker/
2. Установленный Docker Compose: https://docs.docker.com/compose/install/

### Пошаговая инструкция

1. Создайте новую директорию для проекта:
```bash
mkdir url-shortener
cd url-shortener
```

2. Создайте файл .env:
```bash
cat > .env << EOL
STORAGE_TYPE=postgres
PORT=8081
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=urlshortener
DB_SSLMODE=disable
EOL
```

3. Создайте файл docker-compose.yml:
```bash
cat > docker-compose.yml << EOL
version: '3.8'

services:
  app:
    image: marketer7/url-shortener:latest
    container_name: url-shortener
    restart: unless-stopped
    ports:
      - "${PORT:-8081}:8080"
    environment:
      - STORAGE_TYPE=postgres
      - DB_HOST=db
      - DB_PORT=5432
      - DB_USER=\${DB_USER:-postgres}
      - DB_PASSWORD=\${DB_PASSWORD:-your_password}
      - DB_NAME=\${DB_NAME:-urlshortener}
      - DB_SSLMODE=disable
    depends_on:
      - db
    networks:
      - url-shortener-network

  db:
    image: postgres:14-alpine
    container_name: url-shortener-db
    restart: unless-stopped
    environment:
      - POSTGRES_USER=\${DB_USER:-postgres}
      - POSTGRES_PASSWORD=\${DB_PASSWORD:-your_password}
      - POSTGRES_DB=\${DB_NAME:-urlshortener}
    ports:
      - "5433:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - url-shortener-network

networks:
  url-shortener-network:
    driver: bridge

volumes:
  postgres_data:
    driver: local
EOL
```

4. Скачайте и запустите контейнеры:
```bash
# Скачать образ
docker pull marketer7/url-shortener:latest

# Запустить контейнеры
docker-compose up -d
```

5. Проверьте, что всё работает:
```bash
# Проверка статуса контейнеров
docker-compose ps

# Проверка логов
docker-compose logs -f
```

6. Откройте веб-интерфейс:
```
http://localhost:8081/static/test.html
```

### Проверка работоспособности

1. Создание короткой ссылки через curl:
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"original_url":"https://example.com"}' \
  http://localhost:8081/create
```

2. Получение информации о ссылке (замените ABC123 на ваш короткий код):
```bash
curl http://localhost:8081/info/ABC123
```

### Остановка и удаление

```bash
# Остановка контейнеров
docker-compose down

# Удаление всех данных (включая базу)
docker-compose down -v
```

## Быстрый старт с Docker

### Вариант 1: Используя готовый образ из Docker Hub

```bash
# 1. Создайте .env файл из примера
cp .env.example .env

# 2. Запустите приложение через docker-compose
docker-compose up -d

# Сервис будет доступен по адресу http://localhost:8081
```

### Вариант 2: Сборка из исходного кода

```bash
# 1. Клонируйте репозиторий
git clone https://github.com/MarketerKA/shorterApi.git
cd shorterApi

# 2. Создайте .env файл из примера
cp .env.example .env

# 3. Соберите и запустите через docker-compose
docker-compose up --build -d
```

### Проверка работы

```bash
# Проверка статуса контейнеров
docker-compose ps

# Просмотр логов
docker-compose logs -f

# Остановка сервиса
docker-compose down
```

### Переменные окружения

В файле .env можно настроить:
```env
STORAGE_TYPE=postgres    # или memory
PORT=8081               # порт для доступа к API
DB_HOST=db             # для docker-compose
DB_PORT=5432           # внутренний порт PostgreSQL
DB_USER=postgres       # пользователь БД
DB_PASSWORD=your_password
DB_NAME=urlshortener
DB_SSLMODE=disable
```

### Тестирование в Docker

```bash
# Запуск тестов
make docker-test

# Сборка и запуск тестов
make ci
```

## API Endpoints

После запуска доступны следующие endpoints:

- `POST /create` - Создание короткой ссылки
- `GET /{short_code}` - Переход по короткой ссылке
- `GET /info/{short_code}` - Информация о ссылке

### Веб-интерфейс
Доступен по адресу: `http://localhost:8081/static/test.html`

## Разработка

### Локальный запуск без Docker

```bash
# Установка зависимостей
go mod download

# Запуск тестов
make test

# Сборка и запуск
make run
```

### Полезные команды Docker

```bash
# Пересборка образа
docker-compose build

# Перезапуск сервиса
docker-compose restart

# Просмотр логов определенного сервиса
docker-compose logs -f app
docker-compose logs -f db

# Очистка неиспользуемых ресурсов
docker system prune
```

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
```

## Дополнительная информация

- Поддерживаются два типа хранилища: PostgreSQL и In-Memory
- Реализована валидация URL
- Поддержка CORS
- Логирование запросов
- Обработка ошибок
- История запросов в веб-интерфейсе

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
   git clone https://github.com/MarketerKA/shorterApi.git
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









