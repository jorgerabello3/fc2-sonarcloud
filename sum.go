package main

import "fmt"

func main() {
	fmt.Println(sum(10, 5))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func div(a int, b int) int {
	return a / b
}

func times(a int, b int) int {
	return a * b
}

func suX(a int, b int) int {
	return a + b + b
}