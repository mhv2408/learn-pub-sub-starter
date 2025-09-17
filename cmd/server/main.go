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
	rabbitmq_connection_string := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(rabbitmq_connection_string)
	if err != nil {
		log.Fatal("Unable to connect to rabbimq")
		return
	}
	defer conn.Close()
	fmt.Println("Successfully connected to RabbtiMQ!")
	//Signal recognition to exit the program
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Program Shutting down.")

}
