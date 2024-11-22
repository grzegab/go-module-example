package main

import (
	"errors"
	"github.com/grzegab/GO_Module_Example/internal/transport"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (app *Application) EventHandlers(m amqp.Delivery) {
	n, err := getMessageName(m)
	if err != nil {
		m.Reject(false) // No name in header, don't know what to do, reject msg to bin

		return
	}
	ex := transport.NewExchange(app.RabbitDsn)

	switch n {
	case "NewUserWithBasicSchool":
		log.Println(m)
		ev := app.CreateSchool(m.Body)
		ex.Publish(ev)
	default: //No event handler, ignore msg
		log.Printf("[School] No handler found for %s, ignoring\n", n)
	}

	m.Ack(false)
}

func (app *Application) CommandHandlers(m amqp.Delivery) {
	n, err := getMessageName(m)
	if err != nil {
		m.Reject(false) // No name in header, don't know what to do, reject msg to bin

		return
	}

	switch n {
	default:
		log.Printf("No handler found for %s, ignoring\n", n)
		m.Reject(false) // Header has no handler, move to bin, not interested in
	}
}

func getMessageName(m amqp.Delivery) (string, error) {
	mid := m.MessageId
	if mid == "" {
		return "", errors.New("no type found")
	}

	return mid, nil
}
