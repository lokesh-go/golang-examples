package main

import (
	"math"
)

func numRabbits(answers []int) (totalMinRabbits int) {
	hashMap := map[int]int{}
	for _, otherRabbitSameAsHim := range answers {
		hashMap[otherRabbitSameAsHim]++
	}

	for otherRabbitSameAsHim, noOfRabbitThatSay := range hashMap {
		groupOfSameColor := otherRabbitSameAsHim + 1

		noOfGroupNeeds := int(math.Ceil(float64(noOfRabbitThatSay) / float64(groupOfSameColor)))

		noOfRabbits := noOfGroupNeeds * groupOfSameColor

		totalMinRabbits += noOfRabbits
	}

	return totalMinRabbits
}
