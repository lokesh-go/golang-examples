package main

// Rider represents a user who books a ride
type Rider struct {
	id     string
	name   string
	rating float64
}

type riderManagerMethods interface {
	addRider(r Rider)
	getRider(id string) Rider
}

// RiderManager manages riders
type riderManager struct {
	riders map[string]Rider
}

func newRiderManager() riderManagerMethods {
	return &riderManager{
		riders: make(map[string]Rider),
	}
}

func (rm *riderManager) addRider(r Rider) {
	rm.riders[r.id] = r
}

func (rm *riderManager) getRider(id string) Rider {
	return rm.riders[id]
}
