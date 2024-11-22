package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Application struct {
	GrpcAddr        string
	RabbitDsn       string
	DbDsn           string
	ListenPort      string
	CommandQueue    string
	EventQueue      string
	PublishExchange string
}

func NewApplication() *Application {
	return &Application{
		GrpcAddr:     "0.0.0.0:50051",
		RabbitDsn:    "amqp://user:password7@host.docker.internal:5678/",
		DbDsn:        "host=example_database port=5432 user=postgres password=secretpass dbname=school sslmode=disable timezone=UTC connect_timeout=5",
		CommandQueue: "q.example_school.commands",
		EventQueue:   "q.example_school.events",
	}
}

var envLocal = ".env.local"

func (app *Application) UpdateConfig() error {
	envPath := "../../" + envLocal

	if _, err := os.Stat(envPath); err != nil {
		e := fmt.Sprintf("No file found at path: %s", envPath)
		return errors.New(e)
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Println("Error loading .env file")
		return err
	}

	listenPort := os.Getenv("LISTEN_PORT")
	rabbitDsn := os.Getenv("RABBIT_DSN")
	rabbitCommandQueue := os.Getenv("RABBIT_COMMAND_QUEUE")
	rabbitEventQueue := os.Getenv("RABBIT_Event_QUEUE")

	flag.StringVar(&app.ListenPort, "addr", listenPort, "Listen address")
	flag.StringVar(&app.RabbitDsn, "dsn", rabbitDsn, "Rabbit DSN")
	flag.StringVar(&app.CommandQueue, "command", rabbitCommandQueue, "Rabbit Command queue name")
	flag.StringVar(&app.EventQueue, "event", rabbitEventQueue, "Rabbit Event queue name")
	flag.Parse()

	return nil
}
