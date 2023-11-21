package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	arr := strings.Fields(str)

	result := ""

	for i := len(arr) - 1; i >= 0; i-- {
		result += arr[i]
		if i > 0 {
			result += " "
		}
	}
	fmt.Println(result)
}
