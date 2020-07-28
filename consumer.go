package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

const queueName = "TestQueue"

func main() {
	fmt.Println("Consumer application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // TODO: add to env
	if err != nil {
		handleError(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		handleError(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		handleError(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received message: %s\n", d.Body)
		}
	}()

	fmt.Println("Connection successful")
	fmt.Println(" [*] - waiting for messages")
	<-forever
}

func handleError(err error) {
	fmt.Println(err)
	panic(err)
}
