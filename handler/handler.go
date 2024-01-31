package handler

import (
	"net/http"
	"level0/cache"
)

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID заказа из параметров запроса
	orderUID := r.URL.Query().Get("id")
	if orderUID == "" {
		http.Error(w, "Параметр id отсутствует", http.StatusBadRequest)
		return
	}

	// Пытаемся получить данные из кэша
	if orderData, exists := cache.GetOrderFromCache(orderUID); exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(orderData)
	} else {
		// Если данных нет, отправляем статус "Не найдено"
		http.Error(w, "Заказ не найден", http.StatusNotFound)
	}
}
