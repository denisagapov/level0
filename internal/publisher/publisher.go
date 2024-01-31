package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	stan "github.com/nats-io/stan.go"
)

type Message struct {
	Text string `json:"text"`
}

func main() {

	sc, err := stan.Connect("test-cluster", "client-1234")
	if err != nil {
		log.Fatal(err)
	}

	// Пример входящего JSON
	jsonFile, err := os.Open("../../data/payload1.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()
	jsonString, _ := io.ReadAll(jsonFile)
	// Преобразуем JSON в структуру Message
	// fmt.Print(jsonString)
	var message Message
	err = json.Unmarshal([]byte(jsonString), &message)
	if err != nil {
		log.Printf("Ошибка добавления данных %s", err)
	}
	// Публикуем преобразованное сообщение
	err = sc.Publish("foo", []byte(jsonString))
	if err != nil {
		log.Fatal(err)
	}
}
