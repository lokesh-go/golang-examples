package main

type Driver struct {
	id             string
	available      bool
	vehicleDetails vehicle
	rating         float64
}

type DriverManagerMethods interface {
	addDriver(d Driver)
	getDriver(id string) Driver
	getDrivers() []Driver
}

type driverManager struct {
	drivers map[string]Driver
}

func NewDriverManager() DriverManagerMethods {
	return &driverManager{
		drivers: make(map[string]Driver),
	}
}

func (dm *driverManager) addDriver(d Driver) {
	dm.drivers[d.id] = d
}

func (dm *driverManager) getDriver(id string) Driver {
	return dm.drivers[id]
}

func (dm *driverManager) getDrivers() []Driver {
	var drivers []Driver
	for _, d := range dm.drivers {
		drivers = append(drivers, d)
	}
	return drivers
}
