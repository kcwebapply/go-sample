package mq

import (
	"fmt"
	"log"
	"time"

	. "github.com/kcwebapply/go-sample/config"
	"github.com/streadway/amqp"
)

var config Config
var amqpURI string

var queueName string

var Conn *amqp.Connection

func init() {
	// set host strings
	config = GetConfig()
	amqpURI = config.Mq.HOST + ":" + config.Mq.PORT
	queueName = "messageQueue"
}

func Work() {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		fmt.Println("error : ", err)
	}
	channel, _ := conn.Channel()

	messages, _ := consumeMessage(queueName, channel)

	go func() {
		for data := range messages {
			log.Printf("Received message！%s\n", data.Body)
		}
	}()

	go publish(queueName, 10, channel)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")

	forever := make(chan bool)
	<-forever

}

func createQueue(queueName string, channel *amqp.Channel) (amqp.Queue, error) {
	q, err := channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return q, err
}

func consumeMessage(queueName string, channel *amqp.Channel) (<-chan amqp.Delivery, error) {
	messages, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		fmt.Println("Consume error : ", err)
	}
	return messages, err
}

func publish(queueName string, messageNum int, channel *amqp.Channel) {
	for i := 0; i < messageNum; i++ {
		time.Sleep(1 * time.Second)
		var message = fmt.Sprintf("testMessage: %d", i)
		publishMessage(queueName, message, channel)
	}
}
func publishMessage(queueName string, body string, channel *amqp.Channel) error {
	err := channel.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		fmt.Println("publish error : ", err)
	}
	return err
}
