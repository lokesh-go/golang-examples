package main

// Price strategy
type priceStrategy interface {
	calculatePrice(distance float64) (price float64)
}

func newPriceStrategy(userRating float64) priceStrategy {
	switch userRating {
	case 5.0:
		{
			return &ratingBasedPriceStrategy{}
		}
	default:
		return &defaultPriceStrategy{}
	}
}

type defaultPriceStrategy struct{}

func (dp *defaultPriceStrategy) calculatePrice(distance float64) (price float64) {
	rupeesPerKM := 20.0
	return rupeesPerKM * distance
}

type ratingBasedPriceStrategy struct{}

func (rp *ratingBasedPriceStrategy) calculatePrice(distance float64) (price float64) {
	rupeesPerKM := 15.0
	return rupeesPerKM * distance
}

// Driver strategy
type driverFindingStrategy interface {
	findDriver() Driver
}

func newDriverFindingStrategy(dm DriverManagerMethods) driverFindingStrategy {
	return &nearByDriverStrategy{
		drivers: dm.getDrivers(),
	}
}

type nearByDriverStrategy struct {
	drivers []Driver
}

func (nb *nearByDriverStrategy) findDriver() Driver {
	return nb.drivers[0]
}
