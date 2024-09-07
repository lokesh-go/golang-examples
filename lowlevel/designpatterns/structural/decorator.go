package main

import "fmt"

// Adds new functionalities to objects dynamically without altering their structure.
// Decorators provide a flexible alternative to subclassing for extending functionality.
// Example: A Coffee class with a Decorator class like MilkDecorator or SugarDecorator that adds extra features to Coffee

type laptop interface {
	getPrice() int
}

type windowsLaptop struct{}

func (w *windowsLaptop) getPrice() int {
	return 100
}

type laptopWithGraphics struct {
	laptop laptop
}

func (l *laptopWithGraphics) getPrice() int {
	nvidia940MX := 2500
	return l.laptop.getPrice() + nvidia940MX
}

type laptopWith32Ram struct {
	laptop laptop
}

func (l *laptopWith32Ram) getPrice() int {
	ram32GB := 1500
	return l.laptop.getPrice() + ram32GB
}

func main() {
	// Get simple windows laptop
	wL := &windowsLaptop{}
	fmt.Println(wL.getPrice())

	// Now get laptop with graphic card
	lG := &laptopWithGraphics{wL}
	fmt.Println(lG.getPrice())
}
