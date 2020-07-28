package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

const queueName = "TestQueue"

func main() {
	fmt.Println("Go RabbitMQ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // TODO: add to env
	if err != nil {
		handleError(err)
	}
	defer conn.Close()

	fmt.Println("Connection successful")

	ch, err := conn.Channel()
	if err != nil {
		handleError(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		handleError(err)
	}

	fmt.Println(q)

	if err := ch.Publish("", queueName, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte("Hello world")}); err != nil {
		handleError(err)
	}

	fmt.Println("Published message")
}

func handleError(err error) {
	fmt.Println(err)
	panic(err)
}
