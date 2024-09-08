package main

type vehicleType string

const (
	SEDAN vehicleType = "sedan"
	SUV   vehicleType = "suv"
)

type vehicle struct {
	number      string
	driverId    string
	vehicleType vehicleType
}
