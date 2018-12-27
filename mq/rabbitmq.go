package mq

import (
	"fmt"
	"log"

	. "github.com/kcwebapply/go-sample/config"
	"github.com/streadway/amqp"
)

var config Config
var amqpURI string

var Conn *amqp.Connection

func init() {
	// set host strings
	config = GetConfig()
	amqpURI = config.Mq.HOST + ":" + config.Mq.PORT
}

func PublishMessage() {
	conn, err := amqp.Dial(amqpURI)
	fmt.Println("connection error , ", err)
	defer Conn.Close()

	channel, err := conn.Channel()
	fmt.Println("open channel error , ", err)

	q, err := channel.QueueDeclare(
		"messageQueue4", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)

	messages, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)

	fmt.Println("consume error , ", err)

	forever := make(chan bool)

	go func() {
		for data := range messages {
			log.Printf("%s\n", data.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever

}
