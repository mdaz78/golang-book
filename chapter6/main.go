package main

import (
	"fmt"
)

func average(x []float64) float64 {
	total := 0.0

	for _, value := range x {
		total += value
	}

	return total / float64(len(x))
}

// variadic function
func variadicSum(args ...int) int {
	total := 0
	fmt.Println(args)
	for _, v := range args {
		total += v
	}
	return total
}

// closure
func evenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// recursion
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// defer
func one() {
	fmt.Println("one")
}

func two() {
	fmt.Println("two")
}

func three() {
	fmt.Println("three")
}

// panic and recover
func panicAndRecover() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("Panic")
}

// pointer
func zero(xPtr *int) {
	*xPtr = 3
}

/*
Q2 -> Write a function that takes an integer and halves it and returns true if it was even
or false if it was odd. For example, half(1) should return (0, false) and
half(2) should return (1, true).
*/

func questionTwo(x int) (int, bool) {
	value := int(x / 2)
	if value%2 == 0 {
		return value, true
	}
	return value, false
}

func main() {
	// x := []float64{98, 93, 77, 82, 83}
	// y := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(average(x))
	// fmt.Println(variadicSum(1, 2, 3, 4, 5, 6, 7))
	// fmt.Println(variadicSum(y...))
	// nextEven := evenGenerator()
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(factorial(5))

	// defer calls
	// defer three()
	// defer two()
	// one()

	// panic calls
	// panicAndRecover()

	// pointer calls
	// x := 5
	// x := new(int)
	// zero(x)
	// fmt.Println(*x)
	fmt.Println(questionTwo(5))
}
