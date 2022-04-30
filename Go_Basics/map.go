package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61, //Must have trailing comma in multi line
	}

	//number of items
	fmt.Println(len(stocks))

	//Get a value from the map
	fmt.Println(stocks["AMZN"])

	//Get zero value if not found
	fmt.Println(stocks["TSLA"]) // 0

	//Use two value to see if found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	//Set
	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	//Delete
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Single value for is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	fmt.Println("----").
	//Double value "for" is key, value
	for key, value := range stocks {
		fmt.Println(key, ":", value)
	}
}
