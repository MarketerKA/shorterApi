package main

import (
	"log"
	"net/http"

	"url-shortener/internal"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}

	// Инициализация хранилища
	storage, err := internal.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	// Создаем таблицу
	if err := storage.CreateTable(); err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}

	// Инициализация сервиса
	service := internal.NewService(storage)

	// Инициализация обработчика
	handler := internal.NewHandler(service)

	// Настройка маршрутов
	http.HandleFunc("/create", handler.CreateHandler)
	http.HandleFunc("/", handler.GetHandler)

	// Запуск сервера
	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
