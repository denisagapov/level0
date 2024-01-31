package main

import (
	"level0/cache"
	"level0/database"
	"level0/handler"
	"level0/nats"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
)

func main() {

	// Подключение к базе данных PostgreSQL
	db, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Соединение с базой данных установлено")
	}
	defer db.Close()

	// Загрузка кэша
	if err = cache.LoadCacheFromDB(db); err != nil {
		log.Printf("Ошибка при загрузке данных в кэш из базы данных: %s", err)
	}

	// Подключение к NATS
	sc, err := stan.Connect("test-cluster", "client-123")
	if err != nil {
		log.Fatalf("Ошибка при подключении к NATS: %s", err)
	}
	defer sc.Close()

	// Подписка на канал
	subscription, err := nats.SubscribeToChannel(sc, db)
	if err != nil {
		log.Fatalf("Ошибка при подписке на канал NATS: %s", err)
	}
	defer subscription.Unsubscribe()

	// Регистрация обработчика
	http.Handle("/", http.FileServer(http.Dir("../index/")))
	http.HandleFunc("/order", handler.GetOrderHandler)
	log.Println("HTTP сервер запущен с портом 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка при запуске HTTP сервера: %s\n", err)
	}

	select {} // Бесконечный цикл для ожидания сообщений
}
