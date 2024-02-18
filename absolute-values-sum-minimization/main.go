package main

import "fmt"


func absoluteValuesSumMinimization(a []int) int {
	isEven := len(a) % 2 == 0

	if isEven {
		result := a[len(a) / 2 - 1]
		return result
	}

	return a[len(a) / 2]
}


func main() {
	 fmt.Println(absoluteValuesSumMinimization([]int{2, 4, 7}))
	 fmt.Println(absoluteValuesSumMinimization([]int{2, 4, 7, 6})) 
	 fmt.Println(absoluteValuesSumMinimization([]int{2, 4, 7, 6, 6})) 
	 fmt.Println(absoluteValuesSumMinimization([]int{2, 4, 7, 6, 6, 8})) 
}