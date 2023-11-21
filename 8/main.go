package main

import (
	"fmt"
)

func main() {

	var bit int64
	var someNumber int64 = 16

	fmt.Print("input numder bit:\t")
	fmt.Scan(&bit)

	bit = 1 << (bit)

	someNumber = someNumber | bit

	fmt.Printf("%b", someNumber)
}
