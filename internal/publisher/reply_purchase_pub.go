package publisher

import (
	"context"
	"delivery-service/config"
	"delivery-service/internal/domain/message"
	"github.com/gofiber/fiber/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type ReplyPurchasePublisher struct {
	channel *amqp.Channel
	cfg     *config.Config
	conn    *amqp.Connection
}

func NewReplyPurchasePublisher(cfg *config.Config, conn *amqp.Connection) *ReplyPurchasePublisher {
	producer := ReplyPurchasePublisher{
		cfg: cfg,
	}

	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
		return nil
	}
	producer.channel = ch

	return &producer
}

func (pub *ReplyPurchasePublisher) ReplyPurchaseMessage(message *message.CreateOrderReplyMessage) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := ParseOrderToByte(&message)
	if err != nil {
		return err
	}

	log.Infof("Send msgqueue to queue %v - %v",
		pub.cfg.RabbitMQ.CreatePurchaseEvent.Exchange,
		pub.cfg.RabbitMQ.CreatePurchaseEvent.ReplyRoutingKey)

	err = pub.channel.PublishWithContext(ctx,
		pub.cfg.RabbitMQ.CreatePurchaseEvent.Exchange,
		pub.cfg.RabbitMQ.CreatePurchaseEvent.ReplyRoutingKey, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a msgqueue")

	return err
}
