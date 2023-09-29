package vars

import (
	"fmt"
)

func main() {
	// declaring a variable
	var x int = 5 // x := 5 is another way of declaring a variable
	var y int = 7
	var sum int = x + y

	fmt.Println(sum)

	// control statement
	if sum > 10 {
		fmt.Println("The sum is greater than 10")
	} else if sum < 10 {
		fmt.Println("The sum is less than 10")
	} else {
		fmt.Println("The sum is equal to 10")
	}
}
