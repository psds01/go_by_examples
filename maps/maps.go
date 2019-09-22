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
	// does, calling a key not present in the map, change the map?
	fmt.Println("Does it change the map? ->", carToPriceMap)

	// can I "get" with a key of different type?
	// fmt.Println("Getting price of key = 1", carToPriceMap[1])
	// Nah, type error

	priceOfMercedes := carToPriceMap["mercedes"]
	fmt.Println("Price of mercedes:", priceOfMercedes)

	// length of the map
	fmt.Println(len(carToPriceMap))

	// say all BWMs are sold and now we need to remove it
	// syntax : DELETE(mapName, key)
	fmt.Println("Map before deleting a key:", carToPriceMap)
	delete(carToPriceMap, "bmw")
	fmt.Println("Map after deleting a key:", carToPriceMap)

	// what if I try to delete a key not present in the map?
	fmt.Println("Deleting jaguar from the map:", carToPriceMap)
	delete(carToPriceMap, "jaguar")
	fmt.Println("No errors!!", carToPriceMap)

	// get isPresent, value from a map by key
	// syntax: value, isKeyPresent := map[key]
	mercedesPrice, isMercedesAvailble := carToPriceMap["mercedes"]
	fmt.Println("is mercedes availble?", isMercedesAvailble, "price", mercedesPrice)

	jaguarPrice, isJaguarPresent := carToPriceMap["jaguar"]
	if isJaguarPresent {
		fmt.Println("price:", jaguarPrice)
	} else {
		fmt.Println("Jaguar is not available")
	}

}
