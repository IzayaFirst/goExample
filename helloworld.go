package main

import "fmt"

func main() {
	fmt.Println("Input Data :")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println("Input Data is :", input)
	switch input {
	case 0:
		fmt.Println("This is 0")
	case 1:
		fmt.Println("This is 1")
	default:
		fmt.Println("No case")
	}
	/* if input > 10 {
		fmt.Println("More than 10")
	} else {
		fmt.Println("Less than 10 ")
	} */

	// fmt.Println("Hello world")
}
