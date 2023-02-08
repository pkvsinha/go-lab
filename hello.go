package main

import (
	"fmt"
	"time"

	"example/hello/greetings"
)

func timeOfDay() {

}

func main() {
	var hello = "Hello! "
	var time = time.Now()

	fmt.Println(hello, time, greetings.Greet())
}
