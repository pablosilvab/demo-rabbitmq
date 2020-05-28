package main

import (
	"encoding/json"
	"log"

	rabbit "github.com/pablosilvab/demo-rabbitmq"
)

func main() {
	messages := rabbit.ReceiveMsg("test")
	for msg := range messages {
		data, _ := json.Marshal(msg)
		log.Println(string(data))
	}
}
