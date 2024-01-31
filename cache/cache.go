package cache

import (
	"database/sql"
	"sync"
)

var (
	cache   = make(map[string][]byte)
	rwMutex = sync.RWMutex{}
)

// setOrderInCache - Эта функция записывает сериализованные данные о заказе в кэш
func SetOrderInCache(orderUID string, data []byte) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	cache[orderUID] = data
}

// getOrderFromCache - Эта функция пытается получить данные о заказе из кэша
func GetOrderFromCache(orderUID string) ([]byte, bool) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	data, exists := cache[orderUID]
	return data, exists
}

func LoadCacheFromDB(db *sql.DB) error {
	// Выборка всех данных о заказах из БД
	rows, err := db.Query("SELECT order_uid, order_data FROM orders")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var orderUid string
		var orderData []byte
		if err := rows.Scan(&orderUid, &orderData); err != nil {
			return err
		}
		SetOrderInCache(orderUid, orderData)
	}
	return rows.Err()
}
