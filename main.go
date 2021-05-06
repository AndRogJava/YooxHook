package main

import (
	"fmt"
)

// adder function
func adder(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(adder(20, 11))
}
