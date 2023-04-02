package main

import "fmt"

type TestType1 struct {
	Name string
}

type TestType2 struct {
	Name string
}

func (t *TestType1) HelloWorld() string {
	return "Hello World " + t.Name
}

func (t TestType2) HelloWorld() string {
	return "Hello World " + t.Name
}

func HelloWorld[T TestType1 | TestType2](s T) string {
	return fmt.Sprintf("%+v", s)
}

func main() {
	HelloWorld(TestType2{
		"test",
	})
}
