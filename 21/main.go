// Package adapter is an example of the Adapter Pattern.
package main

/*
Реализовать паттерн «адаптер» на любом примере.
*/
import "fmt"

// Target provides an interface with which the system should work.
type Target interface {
	Request() string
}

// Adaptee implements system to be adapted.
type Adaptee struct {
}

// NewAdapter is the Adapter constructor.
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// SpecificRequest implementation.
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Adapter implements Target interface and is an adapter.
type Adapter struct {
	*Adaptee
}

// Request is an adaptive method.
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	//adapter := NewAdapter(&Adaptee{})

	//req := adapter.Request()

	a := Adaptee{}
	fmt.Println(a.SpecificRequest())
	fmt.Println(NewAdapter(&a))

}
