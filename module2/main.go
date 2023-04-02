package main

import (
	"awesomeProject4/module1"
	"log"
)

func main() {
	test := module1.TestType1{
		Name: "test",
	}
	log.Println(test.HelloWorld())
}
