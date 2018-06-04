package main

import (
	"fmt"
)

func main() {
	data := variadicFunc(1, 2, 3, 4, 5, 6, 7, 10)
	fmt.Println(data)
}

func variadicFunc(datas ...int) int {
	total := 0
	for _, round := range datas {
		total += round
	}
	return total
}
