package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func RabbitMqPublisher() {
	fmt.Println("RabbitMQ in starting")

	connection, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	fmt.Println("successfully connected to payload-rabbitmq instance")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()

	// declaring queue with its properties over the channel
	queue, err := channel.QueueDeclare(
		"testing", // queue
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       //args
	)
	if err != nil {
		panic(err)
	}

	body := map[string]any{
		"ProfileID":         251,
		"Msisdn":            2567,
		"Amount":            263.98,
		"ReferenceID":       27573,
		"Reference":         "Samuel",
		"Name":              "some",
		"ProviderRef":       "vodafone",
		"TransactionTypeId": 333333,
		"DepositsId":        1111,
	}

	js, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	js = append(js, '\n')

	// publish message
	err = channel.Publish(
		"",        // exchange
		"testing", // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(js),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Succesfully published message")
}
