package transport

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type EventName int

const (
	NewSchoolCreated EventName = iota
	SchoolDeactivated
	NoHandler
	Error
)

func (m EventName) String() string {
	switch m {
	case NewSchoolCreated:
		return "school.created"
	case SchoolDeactivated:
		return "school.deactivated"
	case NoHandler:
		return "no.handler"
	case Error:
		return "error"
	default:
		return "unknown"
	}
}

type Exchange struct {
	Id   string
	Name string
	Conn *amqp.Connection
	Type EventName
}

func NewExchange(conn *amqp.Connection) *Exchange {
	return &Exchange{
		Name: "ex.events",
		Conn: conn,
	}
}

type EventData struct {
	name EventName
	body []byte
}

func NewEventData(name EventName, body []byte) *EventData {
	return &EventData{
		name: name,
		body: body,
	}
}

func (e *Exchange) Publish(ev *EventData) error {
	ch, closeCh, err := MakeChannel(e.Conn)
	if err != nil {
		log.Printf("Failed to open a channel: %s", err)
		return err
	}
	defer closeCh()

	if err = ch.Publish(e.Name, "", true, false, amqp.Publishing{
		MessageId: e.Id,
		Type:      e.Type.String(),
		Body:      ev.body,
	}); err != nil {
		log.Printf("Failed to publish a message: %s", err)
		return err
	}

	return nil
}
