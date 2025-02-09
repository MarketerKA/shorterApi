package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

// Заменить константу на функцию получения строки подключения
func getDBConnStr() string {
    return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_SSLMODE"))
}

var db *sql.DB

// Структура для хранения URL
type URL struct {
    OriginalURL string `json:"original_url,omitempty"`
    ShortURL    string `json:"short_url,omitempty"`
}

// Создание таблицы в базе данных
func createTable() error {
    query := `
        CREATE TABLE IF NOT EXISTS urls (
            id SERIAL PRIMARY KEY,          -- Уникальный идентификатор
            original_url TEXT NOT NULL,     -- Оригинальная ссылка
            short_url TEXT NOT NULL UNIQUE, -- Короткая ссылка (уникальная)
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Время создания
        )
    `
    _, err := db.Exec(query)
    return err
}

func main() {
    // Загружаем переменные окружения из .env файла
    if err := godotenv.Load(); err != nil {
        log.Fatal("Ошибка загрузки .env файла:", err)
    }

    // Добавьте эти строки перед подключением к БД для отладки
    log.Printf("Connection string: %s", getDBConnStr())
    
    // Подключаемся к базе данных
    var err error
    db, err = sql.Open("postgres", getDBConnStr())
    if err != nil {
        log.Fatal("Ошибка подключения к БД:", err)
    }
    defer db.Close()

    // Проверяем подключение
    err = db.Ping()
    if err != nil {
        log.Fatal("Ошибка проверки подключения к БД:", err)
    }
    log.Println("Успешно подключились к базе данных!")

    // Создаем таблицу
    err = createTable()
    if err != nil {
        log.Fatal("Ошибка создания таблицы:", err)
    }
    log.Println("Таблица успешно создана или уже существует!")
}