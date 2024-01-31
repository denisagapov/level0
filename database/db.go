package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "wildberries"
)

func InitDb() (*sql.DB, error) {
	// Подключение к базе данных PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Установка соединения с базой данных
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	} // Закрытие соединения с базой данных после завершения

	// Проверка соединения с базой данных
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
