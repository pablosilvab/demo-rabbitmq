package rabbit

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

type UpdatesChannel <-chan interface{}

// ReceiveMsg allow receibe a message from a queue. AMQP URL must be defined in the environment variable "RABBIT_URL"
func ReceiveMsg(queue string) UpdatesChannel {
	chs := make(chan interface{})
	var m interface{}

	go func() {
		conn, err := amqp.Dial(os.Getenv("RABBIT_URL"))
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			queue, // name
			false, // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		forever := make(chan bool)
		go func() {
			for d := range msgs {

				Decode(&d.Body, &m)
				chs <- m

			}
		}()
		log.Printf(" [*] Process can be connect to system  Waiting for new request")
		<-forever
	}()

	return chs
}
