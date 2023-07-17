package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"payload-app/api/entities"
	"payload-app/api/models"

	"github.com/streadway/amqp"
)

func RabbitMqConnect() {
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
		"Reference":         "Gemini",
		"Name":              "some",
		"ProviderRef":       "airtel",
		"TransactionTypeId": 728823,
		"DepositsId":        37363,
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

	msgs, err := channel.Consume(
		"testing", // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       //args
	)
	if err != nil {
		panic(err)
	}

	// print consumed messages from queue
	forever := make(chan bool)
	// var payload *entities.IncomingPayload
	go func() {
		for msg := range msgs {
			byteMsg, err := Deserialize(msg.Body)
			if err != nil {
				fmt.Printf("deserialize : %v\n", err)
				// panic(err)
			}

			fmt.Printf("Received Messages: %v\n", byteMsg)

			var incomingPayload entities.IncomingPayload
			err = json.Unmarshal(msg.Body, &incomingPayload)
			if err != nil {
				fmt.Printf("unmarshal: %v\n", err)
				// panic(err)
			}

			fmt.Printf("Received Messages second: %v\n", incomingPayload)

			provider := &entities.CreateProvider{
				DepositsId:  incomingPayload.DepositsId,
				ProviderRef: incomingPayload.ProviderRef,
			}

			var db = DbConnect()
			id, err := models.Insert(provider, db)
			if err != nil {
				fmt.Printf("insert operation: %v\n", err)
				// panic(err)
			}

			fmt.Printf("Database record id: %v\n", id)

		}
	}()

	fmt.Println("waiting for messages...")
	<-forever
}
