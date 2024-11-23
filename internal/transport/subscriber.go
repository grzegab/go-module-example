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
	Conn    *amqp.Connection
	Handler func(m amqp.Delivery)
}

func NewQueue(name string, conn *amqp.Connection, handler HandleFunc) *Queue {
	return &Queue{
		Name:    name,
		Conn:    conn,
		Handler: handler,
	}
}

func (q *Queue) Listen() error {
	var forever chan struct{}
	go func() {
		ch, closeCh, err := MakeChannel(q.Conn)
		if err != nil {
			log.Printf("Failed to open a channel: %s", err)
			return
		}
		defer closeCh()

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
