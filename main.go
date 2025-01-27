package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// RabbitMQ connection URL
	amqpURL := "amqps://student:XYR4yqc.cxh4zug6vje@rabbitmq-exam.rmq3.cloudamqp.com/mxifnklj"

	// Connect to RabbitMQ
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer channel.Close()

	// Declare exchange and queue details
	exchangeName := "exchange.efb0746c-81dd-4124-8728-fed5538418b6"
	queueName := "exam"
	routingKey := "efb0746c-81dd-4124-8728-fed5538418b6"

	// Declare exchange
	err = channel.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		false,        // durable
		true,         // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	// Bind the queue to the exchange
	err = channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to bind the queue: %v", err)
	}

	// Message payload
	message := "Hi CloudAMQP, this was fun!"

	// Publish the message
	err = channel.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: amqp.Persistent, // Make message persistent
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Println("Message sent successfully!")

	// Delete the exchange
	err = channel.ExchangeDelete(
		exchangeName, // name
		false,        // if-unused
		false,        // no-wait
	)
	if err != nil {
		log.Fatalf("Failed to delete the exchange: %v", err)
	}

	log.Println("Exchange deleted successfully!")
}