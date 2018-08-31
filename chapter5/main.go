package main

import (
	"fmt"
)

// array example
func calculateAverage() {
	var total float64
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

// append example
func appendExample() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

// copy example
func copyExample() {
	slice1 := []int{5, 7, 9}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

// map example
func mapExample() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"])

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("I won't be printed")
	}

	if name, ok := elements["H"]; ok {
		fmt.Println(name, ok)
	}

	// nested map
	_elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
	}

	if el, ok := _elements["Be"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

func findSmallest() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
		2, 1, 3, 4,
	}

	// var smallest = x[0]
	// for _, value := range x {
	// 	if value < smallest {
	// 		smallest = value
	// 	}
	// }
	var smallest int

	for index, value := range x {
		if index == 0 || value < smallest {
			smallest = value
		}
	}
	fmt.Println(smallest)
}

func main() {
	// calculateAverage()
	// appendExample()
	// copyExample()
	// mapExample()
	findSmallest()
}
