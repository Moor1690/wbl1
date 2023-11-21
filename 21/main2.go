package main

import "fmt"

// LegacyPrinter представляет старый интерфейс печати
type LegacyPrinter interface {
	Print(s string) string
}

// MyLegacyPrinter старый принтер, который реализует LegacyPrinter
type MyLegacyPrinter struct{}

func (p *MyLegacyPrinter) Print(s string) string {
	newMsg := "Legacy Printer: " + s
	fmt.Println(newMsg)
	return newMsg
}

// NewPrinter представляет новый интерфейс печати
type NewPrinter interface {
	PrintStored() string
}

// MyNewPrinter новый принтер, не совместимый с LegacyPrinter
type MyNewPrinter struct {
	StoredValue string
}

func (p *MyNewPrinter) PrintStored() string {
	fmt.Println(p.StoredValue)
	return p.StoredValue
}

// PrinterAdapter адаптирует MyNewPrinter к LegacyPrinter
type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	NewPrinter NewPrinter
}

func (p *PrinterAdapter) Print(s string) string {
	if p.OldPrinter != nil {
		return p.OldPrinter.Print(s)
	}
	p.NewPrinter.PrintStored()
	return p.NewPrinter.PrintStored()
}

func main() {
	// Старый принтер работает как обычно
	legacyPrinter := &MyLegacyPrinter{}
	legacyPrinter.Print("Hello")

	// Новый принтер требует адаптера
	newPrinter := &MyNewPrinter{StoredValue: "Hello, New Printer!"}
	printerAdapter := &PrinterAdapter{OldPrinter: nil, NewPrinter: newPrinter}
	printerAdapter.Print("Hello")
}
