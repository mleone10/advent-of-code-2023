package day20

import "math"

func FindHouseWithMinPresents(minPresents, maxHousesPerElf, presentsPerHouse int) int {
	houses := map[int]int{}
	for i := 1; i <= minPresents; i++ {
		for j, visits := i, 0; j <= minPresents/presentsPerHouse+presentsPerHouse && visits < maxHousesPerElf; j, visits = j+i, visits+1 {
			houses[j] += i * presentsPerHouse
		}
	}

	min := math.MaxInt
	for h, ps := range houses {
		if ps >= minPresents && h < min {
			min = h
		}
	}

	return min
}
