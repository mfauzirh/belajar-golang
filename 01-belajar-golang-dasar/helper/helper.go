package helper

var version = "1.0.0.0"    // cannot access outside package because start with lower case
var Application = "golang" // can be access outside package becuase start with uppercase

// Cannot be access outsite package
func sayHello2() string {
	return "Hello"
}

func SayHello(name string) string {
	return "Hello " + name
}
