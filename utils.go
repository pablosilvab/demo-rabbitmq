package rabbit

import (
	"encoding/json"
	"log"
)

//Decode This function decode message from mq queue
func Decode(msg *[]byte, m *interface{}) {
	json.Unmarshal(*msg, &m)
	return
}

//Encode This function encode message to send mq queue
func Encode(m interface{}) *[]byte {
	b, err := json.Marshal(m)
	if err != nil {
		log.Println("Error with encode to queue")
	}
	return &b
}
