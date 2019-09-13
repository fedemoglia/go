package maps

import (
	"fmt"
)

func GetMap() map[string]int {
	mapa := map[string]int {
		"Prueba": 3,
	}
	mapa["Fede"] = 1
	mapa["Moglia"] = 2
	delete(mapa, "Fede")
	value, ok := mapa["NO"]
	fmt.Println("Value", value, ok)
	otherValue, key := mapa["Moglia"]
	fmt.Println("Value 2", otherValue, key)
	return mapa
}
