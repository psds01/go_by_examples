// go:maps :: python:dict
// hashes, associative data types

// syntax : mapName := MAKE(MAP[keyType]ValType)

package main

import "fmt"

func main() {
	// declaration syntax = `:=`
	// syntax : mapName := MAKE(MAP[keyType][valType])
	carToPriceMap := make(map[string]int)
	fmt.Println(carToPriceMap)

	// set values
	// syntax: mapName[key] = value
	carToPriceMap["ferrari"] = 300000
	carToPriceMap["lamborghini"] = 250000
	carToPriceMap["bmw"] = 150000
	carToPriceMap["mercedes"] = 125000
	fmt.Println("Prices of cars in USD:", carToPriceMap)

	fmt.Println("Get price of bmw in USD:", carToPriceMap["bmw"])

	// get missing key?
	fmt.Println("Getting price of Out-of-stock car: jaguar", carToPriceMap["jaguar"])

	// can I "get" with a key of different type?
	// fmt.Println("Getting price of key = 1", carToPriceMap[1])
	// Nah, type error
}
