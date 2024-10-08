package binarysearch

/*
	For the bruteforce solution
	Let's say It has started with 1 banana per hour
	and continuously it will try with 1, 2, 3, 4, 5, 6, 7, ......... -> but what is maximum number here ?

	in the piles array we can get maximum numbers
	suppose - piles = [3,6,7,11]
		- get the max number -> 11
		Now we can check one by one from 1 to 11
		1,2,3,4,...........11

	To optimising this we can implement binary search for picking a number from 1,2,3,4,......11 (this is in sorted order)
		- minSpeed is - 1
		- maxSpeed is max number of that array - 11
	and just calculate total hours to each banana
*/

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
