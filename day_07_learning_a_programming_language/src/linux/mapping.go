package main

import "fmt"

func main() {
	// maps are like dictionaries in python
	// built-in make function creates a map
	// make(map [keyType] mapValue)
	shapes := make(map[string]int)

	shapes["triangle"] = 3
	shapes["rectangle"] = 4
	shapes["octagon"] = 8

	fmt.Println(shapes)
}
