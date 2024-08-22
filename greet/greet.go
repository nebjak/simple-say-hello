package greet

import "fmt"

// SayHello returns a greeting message
func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// PrintHello prints a hello message
func PrintHello(name string) {
	fmt.Println(SayHello(name))
}
