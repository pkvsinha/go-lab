package main

import (
	"fmt"
	"time"
)

func interval(fn func(), seconds time.Duration) {
	tick := time.Tick(seconds * time.Millisecond)
	start := time.Now()
	for {
		select {
		case <-tick:
			end := time.Now()
			fmt.Println("Executing the function after ", end.Sub(start))
			fn()
			return
		}
	}
}

func Timeout(fn func(), seconds time.Duration) {
	interval(fn, seconds)
}
