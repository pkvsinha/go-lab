package main

import (
	"fmt"
	"math/rand"
)

func Sum(s []int) int {
	acc := 0
	for _, v := range s {
		acc += v
	}

	return acc
}

func randomInts(n int, c chan int) {
	for i := 0; i <= n; i++ {
		c <- rand.Intn(100)
	}

	close(c)
}

func run(args *[]int, c chan int) {
	c <- Sum(*args)
}

func RunSum() {
	arr := []int{1, 2, 3, 4, 5}
	n := len(arr)
	ch := make(chan int)
	left, right := arr[:n/2], arr[n/2:]
	go run(&left, ch)
	go run(&right, ch)
	lsum, rsum := <-ch, <-ch
	sum := lsum + rsum
	fmt.Println(sum)
}

func RunRandInts() {
	ch := make(chan int, 5)
	go randomInts(10, ch)
	for rnum := range ch {
		fmt.Println(rnum)
	}
}
