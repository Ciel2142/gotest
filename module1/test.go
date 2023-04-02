package module1

import "fmt"

type TestType1 struct {
	Name string
}

type TestType2 struct {
	Name string
}

type HelloWorlder interface {
	HelloWorld() string
}

func (t *TestType1) HelloWorld() string {
	return "Hello World " + t.Name
}

func (t TestType2) HelloWorld() string {
	return "Hello World " + t.Name
}

func HelloWorld(s HelloWorlder) string {
	return fmt.Sprintf("%+v %s", s, s.HelloWorld())
}
