package main

import "fmt"

func main() {
	var conferenceName = "Muffin Conference"
	const conferenceTickets = 27

	// this alternative syntax cannot be used to declare constants
	remainingTickets := 27

	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("There are %v tickets in total, and there are %v tickets still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("You can buy your tickets here to attend.")

	var userName string
	var userTickets int

	// we use a pointer '&' to get the memory address of the variable, that references the actual value
	// Scan will read whatever the user enters and assign the value to the variable
	fmt.Println("Please enter a user name:")
	fmt.Scan(&userName)

	fmt.Println("Please enter how many tickets you want:")
	fmt.Scan(&userTickets)

	fmt.Printf("Welcome %v! You have booked %v tickets. Enjoy!\n", userName, userTickets)
}
