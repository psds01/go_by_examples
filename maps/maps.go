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
}
