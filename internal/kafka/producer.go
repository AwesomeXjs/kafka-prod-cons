package kafka

import (
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"strings"
	"time"
)

// github.com/IBM/sarama
// github.com/segmentio/kafka-go
// github.com/lovoo/goka

var unknownType = errors.New("unknown event type")

const (
	flushTimeout = 5000
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(addresses []string) (*Producer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(addresses, ","),
	}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &Producer{producer: producer}, nil
}

func (p *Producer) Produce(message, topic, key string, timestamp time.Time) error {
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Key:            []byte(key),
		Timestamp:      timestamp,
	}
	// в этот канал мы будем ждать информацию о том что произошло с нашим эвентом (сообщением)
	kafkaChan := make(chan kafka.Event)

	if err := p.producer.Produce(kafkaMessage, kafkaChan); err != nil {
		return err
	}

	event := <-kafkaChan
	switch event.(type) {
	case *kafka.Message:
		m := event.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			return m.TopicPartition.Error
		}
		return nil
	default:
		return unknownType
	}
}

func (p *Producer) Close() {
	p.producer.Flush(flushTimeout) // блокирует завершение до тех пока не завершатся евенты либо пока не сработает таймаут
	p.producer.Close()             // закрываем продюссер
}
