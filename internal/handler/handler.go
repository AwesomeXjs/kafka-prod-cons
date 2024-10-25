package handler

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleMessage(kafkaMsg []byte, offset kafka.Offset, consumerNumber int) error {
	fmt.Println("message:", string(kafkaMsg), "offset:", offset, "consumer number:", consumerNumber)
	return nil
}
