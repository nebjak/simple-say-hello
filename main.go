package simple_say_hello

import (
	"fmt"
	"go.nebjak.dev/simple-say-hello/greet"
)

func PrintHello(name string) {
	fmt.Println(greet.SayHello(name))
}
