package main

import (
	"github.com/grzegab/GO_Module_Example/internal/transport"
	"log"
	"os"
)

func main() {
	log.Println("[School] starting service...")
	app := NewApplication()

	if err := app.UpdateConfig(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("[School] seting up rabbit connection...")
	conn, closeConn, err := transport.MakeAmqpConnection(app.RabbitDsn)
	if err != nil {
		log.Println(err)
		log.Println(err)
		os.Exit(5)
	}
	defer closeConn()
	app.SetUpAmqpConn(conn)

	log.Println("[School] starting listening for rabbit commands calls...")
	cq := transport.NewQueue(app.CommandQueue, app.RabbitConn, app.CommandHandlers)
	if err := cq.Listen(); err != nil {
		log.Println(err)
		os.Exit(2)
	}

	log.Println("[School] starting listening for rabbit event calls...")
	evq := transport.NewQueue(app.EventQueue, app.RabbitConn, app.EventHandlers)
	if err := evq.Listen(); err != nil {
		log.Println(err)
		os.Exit(3)
	}

	log.Println("[School] starting listening for grpc calls (query)...")
	if err := app.Listen(); err != nil {
		log.Println(err)
		os.Exit(4)
	}
}
