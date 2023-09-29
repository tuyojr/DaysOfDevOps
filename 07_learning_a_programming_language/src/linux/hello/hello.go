// Package main: the first line in a go file needs to be the name of the package
// for us to be able to execute it, we need to have a package called main/any other name.
package main

// we can now import different libraries now.
import (
	"fmt"
)

// this is where the program starts
func main() {
	fmt.Println("Wagwan peeps!!!")
	fmt.Println("This is the first Go program.")
}
