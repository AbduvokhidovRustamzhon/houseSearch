package main

import (
	"fmt"
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
		result = sortByExpensivePrice(result)
	}
	return result
}

func searchByMinAndMaxPrice(houses []house, maxPrice int64, minPrice int64) []house { // 3 punkt -
	result := make([]house, 0)
	for _, house := range houses {
		if house.price <= minPrice && house.price >= maxPrice {
			result = append(result, house)
		}
		result = sortByExpensivePrice(result)
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
	houses := []house{
		{
			id:           1,
			name:         "Отличная квартира в центре города",
			price:        45000,
			distanceFrom: 12,
			region:       "Сино",
		},
		{
			id:           2,
			name:         "2х комнатная квартира в Душанбе",
			price:        35000,
			distanceFrom: 32,
			region:       "Сомони",
		},
		{
			id:           3,
			name:         "3х комнатная квартира",
			price:        40000,
			distanceFrom: 52,
			region:       "Ленский",
		},
		{
			id:           4,
			name:         "4х комнатная квартира",
			price:        15000,
			distanceFrom: 32,
			region:       "Сино",
		},
	}

	fmt.Println(sortByCheapPrice(houses))
	fmt.Println(sortByExpensivePrice(houses))
	fmt.Println(sortByFarDistance(houses))
	fmt.Println(sortByCloseDistance(houses))
	fmt.Println(searchByMaxPrice(houses, 42000))
	fmt.Println(searchByMinAndMaxPrice(houses, 17_000, 46_000))
	fmt.Println(searchByRegion(houses, "Sino"))
	fmt.Println(searchByRegions(houses, []string{"Sino", "Somoni"}))
}
