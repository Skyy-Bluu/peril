package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"

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

	channel, err := connection.Channel()

	if err != nil {
		printErrorMsgAndExit("Unable to create channel", err)
	}

	playingState := routing.PlayingState{
		IsPaused: true,
	}

	jsonByte, err := json.Marshal(playingState)

	if err != nil {
		printErrorMsgAndExit("Unable to marshal playing state to JSON", err)
	}

	pubsub.PublishJSON(channel, routing.ExchangePerilDirect, routing.PauseKey, jsonByte)

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
