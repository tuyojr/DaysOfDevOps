package structs

import "fmt"

// a struct is a collection of fields. It's similar to a class in OOP

type person struct {
	name    string
	age     int
	powerUp string
}

func main() {
	// create an object of type person
	u1 := person{"Shisui", 25, "Kotoamatsukami"}
	fmt.Println(u1)

	// use the dot operator to access the fields of the struct
	fmt.Println(u1.name)
}
