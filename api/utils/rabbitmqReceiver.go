package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"payload-app/api/entities"
	"payload-app/api/models"

	"github.com/streadway/amqp"
)

func RabbitMqReceiver() {
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
			var incomingPayload entities.IncomingPayload
			err = json.Unmarshal(msg.Body, &incomingPayload)
			if err != nil {
				fmt.Printf("unmarshal: %v\n", err)
				// panic(err)
			}

			fmt.Printf("Received Messages : %v\n", incomingPayload)

			provider := &entities.CreateProvider{
				DepositsId:  incomingPayload.DepositsId,
				ProviderRef: incomingPayload.ProviderRef,
			}

			//var db = DbConnect()
			db, err := DbConnect()
			if err != nil {
				panic(err)
			}
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
