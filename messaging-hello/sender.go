package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// connecto rabbitmq brocker
// create a queue
// send a message to the queue
// end the program
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

	for i := 0; i < 1000; i++ {
		msg := fmt.Sprintf("%d Einar's message from a broker sender", i)
		err = ch.Publish(
			"einars_football_data", // exchange
			"",                     // routing key
			false,                  // mandatory
			false,                  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			})
		failOnError(err, "Failed to publish a message")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal("%s : %s", msg, err)
	}
}
