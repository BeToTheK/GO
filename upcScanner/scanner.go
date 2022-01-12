package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("A message queue for scanning UPCs")
	fmt.Println("Using http://localhost:15672")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully Connected to RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Test Que",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)
	var scanInput string
	fmt.Print("SCAN UPC:")
	fmt.Scan(&scanInput)

	err = ch.Publish(
		"",
		"Test Que",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(scanInput),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully Published UPC to Queue")
}
