package main

import "fmt"

func main(){
	a := 1
	b := 3
	fmt.Printf("Sum of %v and %v is %v", a, b, add(a, b))
}

func add(a, b int) int {
	return a + b
}