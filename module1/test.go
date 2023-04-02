package module1

type TestType1 struct {
	Name string
}

func (t *TestType1) HelloWorld() string {
	return "Hello World " + t.Name
}
