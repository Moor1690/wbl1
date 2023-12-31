package main

/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/
import (
	"fmt"
	"reflect"
)

func main() {

	var myVar interface{}

	myVar = 42
	myTypeOf(myVar)

	myVar = "Hello, World!"
	myTypeOf(myVar)

	myVar = true
	myTypeOf(myVar)

	myVar = make(chan int)
	myTypeOf(myVar)

	myVar = 42.0
	myTypeOf(myVar)
}

func myTypeOf(val interface{}) {
	fmt.Println(reflect.TypeOf(val))
}
