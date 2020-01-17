package main

import "fmt"

var houses = []house{
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

func ExampleSortByCheapPrice() {

	result := sortByCheapPrice(houses)

	fmt.Println(result)

	//Output:[{4 4х комнатная квартира 15000 32 Сино} {2 2х комнатная квартира в Душанбе 35000 32 Сомони} {3 3х комнатная квартира 40000 52 Ленский} {1 Отличная квартира в центре города 45000 12 Сино}]
}

func ExampleSortByExpensivePrice() {

	result := sortByExpensivePrice(houses)

	fmt.Println(result)

	//Output:[{1 Отличная квартира в центре города 45000 12 Сино} {3 3х комнатная квартира 40000 52 Ленский} {2 2х комнатная квартира в Душанбе 35000 32 Сомони} {4 4х комнатная квартира 15000 32 Сино}]
}

func ExampleSortByNearToCenter() {

	result := sortByCloseDistance(houses)

	fmt.Println(result)

	//Output:[{1 Отличная квартира в центре города 45000 12 Сино} {2 2х комнатная квартира в Душанбе 35000 32 Сомони} {4 4х комнатная квартира 15000 32 Сино} {3 3х комнатная квартира 40000 52 Ленский}]
}

func ExampleSortByFarFromCenter() {

	result := sortByFarDistance(houses)

	fmt.Println(result)

	//Output:[{3 3х комнатная квартира 40000 52 Ленский} {2 2х комнатная квартира в Душанбе 35000 32 Сомони} {4 4х комнатная квартира 15000 32 Сино} {1 Отличная квартира в центре города 45000 12 Сино}]
}

func ExampleSearchByMaxPrice_NoResult() {

	result := searchByMaxPrice(houses, 2000)

	fmt.Println(result)

	//Output:[]

}

func ExampleSearchByMaxPrice_OneResult() {

	result := searchByMaxPrice(houses, 17_000)

	fmt.Println(result)

	//Output:[{4 4х комнатная квартира 15000 32 Сино}]
}

func ExampleSearchByMaxPrice_TwoAndMoreResults() {

	result := searchByMaxPrice(houses, 42_000)

	fmt.Println(result)

	//Output:[{3 3х комнатная квартира 40000 52 Ленский} {2 2х комнатная квартира в Душанбе 35000 32 Сомони} {4 4х комнатная квартира 15000 32 Сино}]
}

func ExampleSearchByLimitPrice_NoResult() {

	result := searchByMinAndMaxPrice(houses, 2_300, 2_600)

	fmt.Println(result)

	//Output:[]

}

func ExampleSearchByLimitPrice_OneResult() {

	result := searchByMinAndMaxPrice(houses, 14_000, 17_000)

	fmt.Println(result)

	//Output:[{4 4х комнатная квартира 15000 32 Сино}]

}

func ExampleSearchByLimitPrice_TwoAndMoreResults() {

	result := searchByMinAndMaxPrice(houses, 17_000, 46_000)

	fmt.Println(result)

	//Output:[{1 Отличная квартира в центре города 45000 12 Сино} {3 3х комнатная квартира 40000 52 Ленский} {2 2х комнатная квартира в Душанбе 35000 32 Сомони}]
}

func ExampleSearchByDistrict_NoResult() {

	result := searchByRegion(houses, "Каленин")

	fmt.Println(result)

	//Output:[]

}

func ExampleSearchByDistrict_OneResult() {

	result := searchByRegion(houses, "Сомони")

	fmt.Println(result)

	//Output:[{2 2х комнатная квартира в Душанбе 35000 32 Сомони}]

}

func ExampleSearchByDistrict_TwoAndMoreResults() {

	result := searchByRegion(houses, "Сино")

	fmt.Println(result)

	//Output:[{1 Отличная квартира в центре города 45000 12 Сино} {4 4х комнатная квартира 15000 32 Сино}]
}

func ExampleSearchByDistricts_NoResult() {

	result := searchByRegions(houses, []string{"Каленин", "Айни"})

	fmt.Println(result)

	//Output:[]

}

func ExampleSearchByDistricts_OneResult() {

	result := searchByRegions(houses, []string{"Сомони", "Каленин"})

	fmt.Println(result)

	//Output:[{2 2х комнатная квартира в Душанбе 35000 32 Сомони}]

}

func ExampleSearchByDistricts_TwoAndMoreResults() {

	result := searchByRegions(houses, []string{"Сино", "Сомони", "Каленин"})

	fmt.Println(result)

	//Output:[{1 Отличная квартира в центре города 45000 12 Сино} {2 2х комнатная квартира в Душанбе 35000 32 Сомони} {4 4х комнатная квартира 15000 32 Сино}]

}
