run-p:
	go run cmd/producer/main.go

run-c:
	go run cmd/consumer/main.go

compose:
	docker-compose up -d

get-deps:
	go get -u github.com/confluentinc/confluent-kafka-go/v2/kafka