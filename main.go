package main

import (
	"log"
	"net/http"
	"os"

	"url-shortener/internal"
	"url-shortener/internal/storage"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}

	// Инициализация хранилища
	store, err := storage.NewStorage(os.Getenv("STORAGE_TYPE"))
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	// Создаем таблицу
	if err := store.CreateTable(); err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}

	// Инициализация сервиса
	service := internal.NewService(store)

	// Инициализация обработчика
	handler := internal.NewHandler(service)

	// Добавляем обработку статических файлов
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Настройка маршрутов
	http.HandleFunc("/create", handler.CreateHandler)
	http.HandleFunc("/info/", handler.InfoHandler)
	http.HandleFunc("/", handler.GetHandler)

	// Запуск сервера
	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
