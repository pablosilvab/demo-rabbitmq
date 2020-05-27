package rabbit

import (
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func SendMsg(queueName string, message interface{}) error {
	conn, err := amqp.Dial(os.Getenv("RABBIT_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	msg, err := json.Marshal(message)

	ch, err := conn.Channel()
	err = failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	err = failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         msg,
		})
	failOnError(err, "Failed to publish a message")

	//	elastic.Log("rabbit", Log{q.Name, time.Now(), "send", "OK", ""})

	log.Printf(" [x] Sent %+v", message)
	return nil
}

func failOnError(err error, msg string) error {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		return err
	}
	return nil
}
