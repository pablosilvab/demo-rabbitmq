package main

import (
	"fmt"

	rabbit "github.com/pablosilvab/demo-rabbitmq"
)

func main() {
	msg := "Mensaje de prueba"
	fmt.Print(msg)
	rabbit.SendMsg("test", msg)
}
