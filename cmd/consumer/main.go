package main

import (
	"github.com/AwesomeXjs/kafka-prod-cons/internal/handler"
	"github.com/AwesomeXjs/kafka-prod-cons/internal/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	addresses     = []string{"localhost:9091", "localhost:9092", "localhost:9093"}
	topic         = "my-topic"
	consumerGroup = "my-group"
)

func main() {
	h := handler.NewHandler()
	c1, err := kafka.NewConsumer(h, addresses, topic, consumerGroup, 1)
	if err != nil {
		log.Fatal(err)
	}
	c2, err := kafka.NewConsumer(h, addresses, topic, consumerGroup, 2)
	if err != nil {
		log.Fatal(err)
	}

	c3, err := kafka.NewConsumer(h, addresses, topic, consumerGroup, 3)
	if err != nil {
		log.Fatal(err)
	}

	go func() {

		c1.Start()
	}()

	go func() {

		c2.Start()
	}()

	go func() {

		c3.Start()
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	if err = c1.Stop(); err != nil {
		log.Fatal(err)
	}
	if err = c2.Stop(); err != nil {
		log.Fatal(err)
	}
	if err = c3.Stop(); err != nil {
		log.Fatal(err)
	}
	
}
