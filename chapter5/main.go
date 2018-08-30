package main

import (
	"fmt"
)

// array example
func calculateAverage() {
	var total float64 = 0
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	// for i := 1; i < len(x); i++ {
	// 	total += float64(x[i])
	// }
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func main() {
	calculateAverage()
}
