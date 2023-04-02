package main

import (
	"fmt"
	"github.com/Ciel2142/gotest/awesomeProject4/module1"
)

func main() {
	x := module1.TestType1{
		Name: "test",
	}

	fmt.Println(x.HelloWorld())
	s := module1.HelloWorld(&x)
	fmt.Println(s)
}
