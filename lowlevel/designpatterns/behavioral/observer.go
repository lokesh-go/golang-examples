package main

import "fmt"

// Behavioral design patterns focus on how objects communicate and interact with each other.

// Observer Establishes a one-to-many relationship between objects
// where a state change in one object notifies and updates all its dependents (observers)

type observer interface {
	update(temperature float64)
}

// Have to notify all observers when temperature has updated
type weatherStation struct {
	observers   []observer
	temperature float64
}

// Observer operations
func (w *weatherStation) addObserver(obj observer) {
	w.observers = append(w.observers, obj)
}

func (w *weatherStation) removeObserver(obj observer) {
	for i, observer := range w.observers {
		if observer == obj {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *weatherStation) notifyAll() {
	for _, observer := range w.observers {
		observer.update(w.temperature)
	}
}

func (w *weatherStation) setTemperature(temperature float64) {
	fmt.Println("Setting temperature to ", temperature)
	w.temperature = temperature
	w.notifyAll()
}

// Observers
type iPhoneDisplay struct{}

func (p *iPhoneDisplay) update(temperature float64) {
	fmt.Println("Updated temperature at Phone:", temperature)
}

type iPadDisplay struct{}

func (p *iPadDisplay) update(temperature float64) {
	fmt.Println("Updated temperature at iPad:", temperature)
}

func main() {
	ws := &weatherStation{}

	// Observers
	iPhone := &iPhoneDisplay{}
	iPad := &iPadDisplay{}
	ws.addObserver(iPhone)
	ws.addObserver(iPad)

	ws.setTemperature(30.123)
	ws.removeObserver(iPhone)
	ws.setTemperature(40.023)
}
