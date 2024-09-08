package main

import "fmt"

type tripStatus string

const (
	started    tripStatus = "started"
	completed  tripStatus = "completed"
	inprogress tripStatus = "inprogress"
)

type location struct {
	lat  float64
	long float64
}

type trip struct {
	id      string
	rider   *Rider
	driver  *Driver
	srcLoc  location
	destLoc location
	status  tripStatus
	price   float64
}

type tripManagerMethods interface {
	createTrip(r *Rider, srcLoc, dstLoc location, dm DriverManagerMethods) (t *trip)
	displayTrip()
}

// TripManager manages trips
type tripManager struct {
	trips map[string]trip
}

func newTripManager() tripManagerMethods {
	return &tripManager{
		trips: make(map[string]trip),
	}
}

func (m *tripManager) createTrip(r *Rider, srcLoc, dstLoc location, dm DriverManagerMethods) (t *trip) {
	ds := newDriverFindingStrategy(dm)
	driver := ds.findDriver()

	ps := newPriceStrategy(r.rating)
	price := ps.calculatePrice(calDistance(srcLoc, dstLoc))

	t = &trip{
		id:      "T01",
		rider:   r,
		driver:  &driver,
		srcLoc:  srcLoc,
		destLoc: dstLoc,
		status:  started,
		price:   price,
	}
	m.trips[t.id] = *t
	return t
}

func (m *tripManager) displayTrip() {
	for _, t := range m.trips {
		fmt.Println(t)
	}
}

func calDistance(srcLoc, dstLoc location) float64 {
	return 15.0
}
