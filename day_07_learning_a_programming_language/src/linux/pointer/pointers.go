package pointer

import (
	"fmt"
)

func main() {
	num := 10
	// we reference the variable by using the & operator
	// this gives us a pointer to the variable in memory
	incrementer(&num)
	fmt.Println("Value of num is: ", num)

	// & operator returns the address of the variable
	// this gives us a pointer to the variable in memory
	// this is called referencing
	fmt.Println("Address of num is: ", &num)
}

// creating a function to increment the value of a variable
// this function takes a pointer to an integer as an argument
// this is called dereferencing
func incrementer(val *int) {
	// we use the * operator to dereference the pointer
	*val++
}
