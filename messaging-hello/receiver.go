package main

import (
	"log"

	"github.com/streadway/amqp"
)

// connecto rabbitmq brocker
// connect to a queue
// listen for messaging incoming to this queue
// program waiting for more messaging progrmas to come
func main() {
	conn, err := amqp.Dial("amqp://admin:Password123@159.65.220.217:5672")
	failOnError(err, "Failed to connect")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"einars_football_data", // name
		"fanout",               // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	queue, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to create queue")

	err = ch.QueueBind(
		queue.Name,             // queue name
		"",                     // routing key
		"einars_football_data", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to consume queue messages")

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)
			d.Ack(true)
		}
	}()
	log.Print("[*] Waiting for messages, please exit the program to stop")
	forever := make(chan bool)
	<-forever

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal("%s : %s", msg, err)
	}
}
