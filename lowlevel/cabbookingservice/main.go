package main

import "fmt"

func main() {
	// Creates Rider and Driver
	rider1 := Rider{"R01", "Rider 1", 5.0}
	rider2 := Rider{"R02", "Rider 2", 4.0}
	driver1 := Driver{"D01", true, vehicle{"KA01", "D01", SEDAN}, 4.0}

	// Rider
	rm := newRiderManager()
	rm.addRider(rider1)
	rm.addRider(rider2)

	// Driver
	dm := NewDriverManager()
	dm.addDriver(driver1)

	// Trip
	tm := newTripManager()
	t := tm.createTrip(&rider1, location{10.0, 17.0}, location{100.0, 200.0}, dm)
	tm.displayTrip()

	// Payment
	pf := newPaymentFactory(cash)
	fmt.Println(pf.pay(t.price))
}
