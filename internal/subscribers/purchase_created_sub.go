package subscribers

import (
	"context"
	"delivery-service/config"
	"delivery-service/internal/domain/message"
	"delivery-service/internal/service/packageserv"
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

type PurchaseCreatedSub struct {
	cfg         *config.Config
	conn        *amqp.Connection
	packageServ *packageserv.ShippingPackageService
}

func NewPurchaseCreatedSub(cfg *config.Config, conn *amqp.Connection, packageServ *packageserv.ShippingPackageService) *PurchaseCreatedSub {
	return &PurchaseCreatedSub{
		cfg:         cfg,
		conn:        conn,
		packageServ: packageServ,
	}
}

func (pb *PurchaseCreatedSub) ListenPurchaseEvent(wg *sync.WaitGroup) {
	channel, err := pb.conn.Channel()
	defer channel.Close()

	// define an exchange type "topic"
	err = channel.ExchangeDeclare(
		pb.cfg.RabbitMQ.CreatePurchaseEvent.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("cannot declare exchange: %v", err)
	}

	// create queue
	q, err := channel.QueueDeclare(
		"purchase_delivery_commit",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("cannot declare queue: %v", err)
	}

	err = channel.QueueBind(
		q.Name,
		pb.cfg.RabbitMQ.CreatePurchaseEvent.CommitRoutingKey,
		pb.cfg.RabbitMQ.CreatePurchaseEvent.Exchange,
		false,
		nil)
	if err != nil {
		log.Fatalf("cannot bind exchange: %v", err)
	}

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		q.Name,                      // queue
		pb.cfg.RabbitMQ.ServiceName, // consumer
		true,                        // auto ack
		false,                       // exclusive
		false,                       // no local
		false,                       // no wait
		nil,                         //args
	)
	if err != nil {
		panic(err)
	}

	defer wg.Done()
	// handle consumed messages from queue
	for msg := range msgs {
		log.Infof("received message from: %s", msg.RoutingKey)
		if err := pb.handleMessage(msg); err != nil {
			log.Infof("The handling message was failed cause %s", err)
		}

	}

	log.Infof("message queue has started")
	log.Infof("waiting for messages...")
}

func (pb *PurchaseCreatedSub) handleMessage(msg amqp.Delivery) error {
	startTime := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msgDTO := message.CreateOrderMessage{}

	if err := json.Unmarshal(msg.Body, &msgDTO); err != nil {
		log.Infof("Parse message to order failed cause: %s", err)
		return err
	}

	if err := pb.packageServ.CreateShippingPackage(ctx, &msgDTO); err != nil {
		log.Infof("Handling order message was failed cause: %s", err)
		return err
	}

	endTime := time.Now()
	log.Infof("The order [%v]  was processed successfully - duration:%v", msgDTO.OrderID, endTime.Sub(startTime))
	return nil
}
