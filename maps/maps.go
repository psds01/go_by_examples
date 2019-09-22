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

}
