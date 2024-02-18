package main

import "fmt"

func add1(a, b int) int {
	return a + b
}

func add2(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println(add2(1, 2))
	fmt.Println(add2(2, 4, 6, 8))
}