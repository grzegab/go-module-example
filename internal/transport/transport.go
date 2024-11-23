package transport

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func MakeAmqpConnection(dns string) (*amqp.Connection, func(), error) {
	conn, err := amqp.Dial(dns)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err)
		return nil, nil, nil
	}

	defClose := func() {
		err := conn.Close()
		if err != nil {
			log.Printf("Failed to close RabbitMQ connection: %s", err)
		}
	}

	return conn, defClose, err
}

func MakeChannel(conn *amqp.Connection) (*amqp.Channel, func(), error) {
	ch, err := conn.Channel() // don't share channels between threads, each thread one channel
	if err != nil {
		log.Printf("Failed to open a channel: %s", err)
		return nil, nil, err
	}

	closeChannel := func() {
		err := ch.Close()
		if err != nil {
			log.Printf("Failed to close a channel: %s", err)
		}
	}

	return ch, closeChannel, err
}
