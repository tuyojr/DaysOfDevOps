package main

import "fmt"

func main() {
	// the elements of an array can be of fixed type and size
	fixedArr := [5]int{1, 2, 3, 4, 5}

	// the elements of a slice can be of any type and size
	slice := fixedArr[:]
	slice = append(slice, 6)

	dynamicArr := []int{1, 2, 3, 4, 5}
	dynamicArr = append(dynamicArr, 6, 7, 8, 9, 10)

	fmt.Println(fixedArr)
	fmt.Println(slice)
	fmt.Println(dynamicArr)
}
