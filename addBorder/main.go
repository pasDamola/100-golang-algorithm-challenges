package main

import "fmt"

func addBorder(a []string)  {
	wallLength := len(a) + 2
	border := ""

	for i := 0; i < wallLength; i++ {
		border += "*"
	}

	a = append(a, border)
	newSlice := append([]string{border}, a...)

	fmt.Println(newSlice)

}


func main() {
	addBorder([]string{"abc", "cde"}) // ["*****", "*abc*", "*cde*", "*****"]
}