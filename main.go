package main

import "fmt"

func main() {

	fmt.Println(hello())
	fmt.Println(plus(2, 3))
}

func hello() string {
	return "Hello HokaSonja"
}

func plus(a, b int) int {
	return a + b
}
