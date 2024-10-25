package main

import (
	"fmt"
	"github.com/AwesomeXjs/kafka-prod-cons/internal/kafka"
	"github.com/google/uuid"
	"log"
	"time"
)

const (
	topic        = "my-topic"
	numberOfKeys = 20
)

var addresses = []string{"localhost:9091", "localhost:9092", "localhost:9093"}

func main() {
	// регистрируем продюсер
	producer, err := kafka.NewProducer(addresses)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	keys := generateUUIDstring()

	for i := 0; i < 1000; i++ {
		msg := fmt.Sprintf("message %d", i)
		key := keys[i%numberOfKeys]

		// отправляем сообщение в кафку
		if err = producer.Produce(msg, topic, key, time.Now()); err != nil {
			log.Fatal(err)
		}
	}

}

func generateUUIDstring() [numberOfKeys]string {
	var uuids [numberOfKeys]string
	for i := 0; i < numberOfKeys; i++ {
		uuids[i] = uuid.NewString()
	}
	return uuids
}
