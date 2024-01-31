package nats

import (
	"database/sql"
	"encoding/json"
	"log"
	"level0/cache"
	"level0/internal/model"
	"level0/validation"

	"github.com/nats-io/stan.go"
)

// Подписка на канал "foo" с обработчиком сообщений
func SubscribeToChannel(sc stan.Conn, db *sql.DB) (stan.Subscription, error) {
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {

		// Десериализация в структуру Order

		var order model.Order

		if err := json.Unmarshal(m.Data, &order); err != nil {
			log.Printf("Ошибка при десериализации JSON: %s\n", err)
			return

		}

		if err := validation.ValidateOrder(order); err != nil {
			log.Printf("Ошибка валидации: %s\n", err)
			return

		}

		// Обработка заказа

		orderData, err := json.Marshal(order)
		if err != nil {
			log.Println("Не удалось сериализовать order:", err)
			return

		}

		// Проверка кэша и запись в него
		if _, exists := cache.GetOrderFromCache(order.OrderUID); !exists {
			cache.SetOrderInCache(order.OrderUID, orderData)

			// Если нет в кэше, то вставляем данные в БД

			insertQuery := `INSERT INTO orders (order_uid, order_data) VALUES ($1, $2)`
			if _, err = db.Exec(insertQuery, order.OrderUID, orderData); err != nil {
				log.Printf("Ошибка при записи в базу данных: %s\n", err)
			} else {
				log.Printf("Добавлен новый заказ '%s'", order.OrderUID)
			}
		} else {
			// Если данные уже есть в кэше

			log.Printf("Данные для orderUID '%s' найдены в кэше, вставка в БД пропущена.", order.OrderUID)
		}
	})
	if err != nil {
		log.Fatalf("Ошибка при подписке на канал NATS: %s\n", err)
	}
	return sub, err
}
