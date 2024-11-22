package transport

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type TransportListener interface {
	Listen() error
}

type HandleFunc func(m amqp.Delivery)

type Queue struct {
	Name    string
	Dsn     string
	Handler func(m amqp.Delivery)
}

func NewQueue(name, dsn string, handler HandleFunc) *Queue {
	return &Queue{
		Name:    name,
		Dsn:     dsn,
		Handler: handler,
	}
}

func (q *Queue) Listen() error {
	conn, err := amqp.Dial(q.Dsn)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err)
		return err
	}
	defer conn.Close()

	var forever chan struct{}
	go func() {
		ch, err := conn.Channel() // don't share channels between threads, each thread one channel
		if err != nil {
			log.Printf("Failed to open a channel: %s", err)
			return
		}
		defer ch.Close()

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		if err != nil {
			log.Printf("Failed to register a consumer: %s", err)
			return
		}

		for m := range msgs {
			q.Handler(m)
		}
	}()
	<-forever

	return nil
}
