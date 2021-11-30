package main

import (
	"fmt"
)

func main() {
	var fruit string = "6 apples"
	var watermallan bool = false
	if watermallan {
		fruit = "1 apple"
	} else {
		fruit = "6 apples"
	}
	fmt.Println("buy: ", fruit)
}
