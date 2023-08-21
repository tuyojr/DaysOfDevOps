package main

import "fmt"

func main() {
	// the elements of an array can be of fixed type and size
	fixed_arr := [5]int{1, 2, 3, 4, 5}

	// the elements of a slice can be of any type and size
	slice := fixed_arr[:]
	slice = append(slice, 6)

	dynamic_arr := []int{1, 2, 3, 4, 5}
	dynamic_arr = append(dynamic_arr, 6, 7, 8, 9, 10)

	fmt.Println(fixed_arr)
	fmt.Println(slice)
	fmt.Println(dynamic_arr)
}
