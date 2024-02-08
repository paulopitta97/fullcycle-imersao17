package main

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/paulopitta97/fullcycle-imersao17/payment/internal/entity"
	"github.com/paulopitta97/fullcycle-imersao17/payment/pkg/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

// {"order_id":"123","card_hash":"123","total":100.00}
func main() {
	ctx := context.Background()
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp091.Delivery)
	// rodando em background / outra thread
	go rabbitmq.Consume(ch, msgs, "orders")

	for msg := range msgs {
		var orderRequest entity.OrderRequest
		err := json.Unmarshal(msg.Body, &orderRequest)

		if err != nil {
			slog.Error(err.Error())
			break
		}

		response, err := orderRequest.Process()
		if err != nil {
			slog.Error(err.Error())
			break
		}

		responseJson, err := json.Marshal(response)
		if err != nil {
			slog.Error(err.Error())
			break
		}

		err = rabbitmq.Publish(ctx, ch, string(responseJson), "amq.direct")
		if err != nil {
			slog.Error(err.Error())
			break
		}

		msg.Ack(false)
		slog.Info("Order processed")
	}
}
