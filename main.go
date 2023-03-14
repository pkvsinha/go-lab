package main

import "fmt"

func main() {
	RunSum()
	RunRandInts()
	fn := func() {
		fmt.Println("Hello !")
	}
	Timeout(fn, 60*1000)
	// StartServer()
}
