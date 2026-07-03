package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	rabbitMQServerURL := "amqp://guest:guest@localhost:5672/"

	connection, err := amqp.Dial(rabbitMQServerURL)

	if err != nil {
		printErrorMsgAndExit("Error establishing connection to RabbitMQ server", err)
	}

	defer connection.Close()

	fmt.Println("Connection to RabbitMQ server successful")

	runUntilUserExits()
}

func printErrorMsgAndExit(contextMsg string, err error) {
	fatalErr := fmt.Errorf("%s: %v", contextMsg, err)

	log.Fatal(fatalErr)
}

func runUntilUserExits() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
