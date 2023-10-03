package main

import "fmt"

func main() {
	var conferenceName = "Muffin Conference"
	const conferenceTickets = 27

	// this alternative syntax cannot be used to declare constants ==> variable := value
	var remainingTickets uint = 27

	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("There are %v tickets in total, and there are %v tickets still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("You can buy your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// we use a pointer '&' to get the memory address of the variable, that references the actual value
	// Scan will read whatever the user enters and assign the value to the variable
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	fmt.Println("Please enter how many tickets you want:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v! You have successfully bought %v tickets, you should get a confirmation shortly in your email: %v.\n", firstName, lastName, userTickets, email)

	fmt.Printf("There are now %v tickets left.\n", remainingTickets)
}
