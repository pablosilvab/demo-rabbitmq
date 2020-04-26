package main

import (
	"fmt"

	"github.com/pablosilvab/demo-rabbitmq/sender"
)

func main() {
	msg := "Mensaje de prueba"
	fmt.Print(msg)
	sender.SendMsg("test", msg)
}
