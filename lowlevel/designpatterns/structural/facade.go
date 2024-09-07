package main

import "fmt"

// It is a structural design pattern that provides a simplified interface to a complex subsystem
// It hides the complexity of the system by exposing a unified, easier-to-use interface
// This is helpful when you have multiple classes or a complex system and want to reduce the learning curve for the clients.

type mapVendor interface {
	GetGeoCodeDetails() string
}

type googleMap struct{}

func (g *googleMap) GetGeoCodeDetails() string {
	return "Google Maps"
}

type mapMyIndia struct{}

func (m *mapMyIndia) GetGeoCodeDetails() string {
	return "Map My India"
}

type hereMap struct{}

func (h *hereMap) GetGeoCodeDetails() string {
	return "Here Maps"
}

// MapFacade ...
type MapFacade struct {
	googleMap  mapVendor
	hereMap    mapVendor
	mapMyIndia mapVendor
}

// New map facade
func newMapFacade() *MapFacade {
	return &MapFacade{
		googleMap:  &googleMap{},
		hereMap:    &hereMap{},
		mapMyIndia: &mapMyIndia{},
	}
}

func main() {
	mapFacade := newMapFacade()
	fmt.Println(mapFacade.hereMap.GetGeoCodeDetails())
	fmt.Println(mapFacade.mapMyIndia.GetGeoCodeDetails())
	fmt.Println(mapFacade.googleMap.GetGeoCodeDetails())
}
