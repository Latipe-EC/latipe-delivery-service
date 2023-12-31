package message

import (
	"context"
	"delivery-service/config"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderWorkerProducer struct {
	channel *amqp.Channel
	cfg     *config.Config
}

var producer OrderWorkerProducer

func InitWorkerProducer(config *config.Config) error {
	conn, err := amqp.Dial(config.RabbitMQ.Connection)
	failOnError(err, "Failed to connect to RabbitMQ")
	log.Printf("[%s] Producer has been connected", "INFO")

	producer.cfg = config
	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
		return err
	}
	producer.channel = ch

	failOnError(err, "Failed to declare a queue")

	return nil
}

func SendEmailMessage(request interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := ParseToMessage(&request)
	if err != nil {
		return err
	}

	err = producer.channel.ExchangeDeclare(
		producer.cfg.RabbitMQ.EmailEvent.Exchange, // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)

	log.Printf("[Info] [%s] Send message %v", producer.cfg.RabbitMQ.EmailEvent.RoutingKey, request)
	err = producer.channel.PublishWithContext(ctx,
		producer.cfg.RabbitMQ.EmailEvent.Exchange,   // exchange
		producer.cfg.RabbitMQ.EmailEvent.RoutingKey, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a email message")

	return nil
}
