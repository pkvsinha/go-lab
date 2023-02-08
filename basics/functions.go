package basics

import "fmt"

func Swap(a, b int) (x, y int) {
	x = b
	y = a
	return
}

func Greet(a, b string, times int) string {
	return a + b + fmt.Sprint(times)
}
