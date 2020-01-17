package main

import (
	"sort"
)

type house struct {
	id           int64
	name         string
	price        int64
	distanceFrom int64
	region       string
}

func sortBy(houses []house, comparing func(a, b house) bool) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return comparing(result[i], result[j])
	})
	return result
}

func sortByCheapPrice(houses []house) []house { // 1 punkt A
	return sortBy(houses, func(a, b house) bool { return a.price < b.price })
}

func sortByExpensivePrice(houses []house) []house { // 1 punkt B
	return sortBy(houses, func(a, b house) bool { return a.price > b.price })
}

func sortByFarDistance(houses []house) []house { // 2 punkt A
	return sortBy(houses, func(a, b house) bool { return a.distanceFrom > b.distanceFrom })
}

func sortByCloseDistance(houses []house) []house { // 2 punkt B
	return sortBy(houses, func(a, b house) bool { return a.distanceFrom < b.distanceFrom })

}

func searchByMaxPrice(houses []house, maxPrice int64) []house { // 3 punkt -
	result := make([]house, 0)
	copy(result, houses)
	for _, house := range houses {
		if house.price <= maxPrice {
			result = append(result, house)
		}
	}
	return result
}

func searchByMinAndMaxPrice(houses []house, maxPrice int64, minPrice int64) []house { // 3 punkt -
	result := make([]house, 0)
	for _, house := range houses {
		if house.price <= minPrice && house.price >= maxPrice {
			result = append(result, house)
		}
	}
	return result
}

func searchByRegion(houses []house, region string) []house { // 3 punkt -
	result := make([]house, 0)
	for _, house := range houses {
			if house.region == region {
				result = append(result, house)
			}
	}
	return result
}


func searchByRegions(houses []house, regions []string) []house { // 3 punkt -
	result := make([]house, 0)
	for _, house := range houses {
		for _, region := range regions {
			if house.region == region {
				result = append(result, house)
			}
		}
	}
	return result
}

func main() {

}
