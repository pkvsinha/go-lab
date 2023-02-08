package basics

import (
	"runtime"
	"time"
)

func loops() {
	var sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	i := 0
	sumx := 0
	for i < 10 {
		// acts as a while loop
		sumx += i
		i++
	}
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}

func os() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "Mac"
	case "linux":
		return "Linux"
	default:
		return os
	}
}

func timeOfDay() string {
	t := time.Now().Hour()
	switch {
	case t < 12:
		return "Morning"
	case t == 12:
		return "Noon"
	case t > 12:
		return "Afternoon"
	default:
		return "Evening"
	}
}
