package main

import "fmt"

func main() {
	// the elements of an array can be of fixed type and size
	fixedArr := [5]int{1, 2, 3, 4, 5}

	// the elements of a slice can be of any type and size
	// creating a slice from an array
	slice := fixedArr[:]
	slice = append(slice, 6)

	// creating a dynamic array with the built-in append() function
	dynamicArr := []int{1, 2, 3, 4, 5}
	dynamicArr = append(dynamicArr, 6, 7, 8, 9, 10)

	// dynamic string array
	strArr := []string{"Shisui", "Madara", "Obito", "Itachi"}
	strArr = append(strArr, "Sasuke", "Fugaku")

	fmt.Println(fixedArr)
	fmt.Println(slice)
	fmt.Println(dynamicArr)
}
