package main

import (
	rabbit "github.com/pablosilvab/demo-rabbitmq"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "Pablo",
		Age:  21,
	}

	rabbit.SendMsg("test", person)
}
