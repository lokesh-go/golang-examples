package binarysearch

import "math"

func minEatingSpeed(piles []int, h int) int {
	high := getMaxKToEatBanana(piles)
	low := 1

	for low <= high {
		mid := (low + high) / 2

		// Calculate total hour to eat banana it eating per hour is mid
		totalHourToEatBanana := calTotalHourToEatBanana(piles, mid)
		if totalHourToEatBanana <= h {
			// We can also store mid and further calculated minimum of possible answers
			// Or Just return the "low"
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func getMaxKToEatBanana(piles []int) int {
	m := math.MinInt
	for _, num := range piles {
		m = max(m, num)
	}
	return m
}

func calTotalHourToEatBanana(piles []int, eatingPerHour int) (totalHourToEatBanana int) {
	for _, num := range piles {
		hourToEatBananaPerPile := int(math.Ceil(float64(num) / float64(eatingPerHour)))
		totalHourToEatBanana += hourToEatBananaPerPile
	}
	return totalHourToEatBanana
}
