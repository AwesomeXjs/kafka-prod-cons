run-p:
	go run cmd/producer/main.go

compose:
	docker-compose up -d

get-deps:
	go get -u github.com/confluentinc/confluent-kafka-go/v2/kafka