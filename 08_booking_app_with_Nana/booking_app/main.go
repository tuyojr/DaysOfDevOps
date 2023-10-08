package main

import (
	"fmt"
	"strings"
)

func main() {
	// this alternative syntax cannot be used to declare constants ==> variable := value
	conferenceName := "Muffin Conference"
	const conferenceTickets = 27

	var remainingTickets uint = 27
	var bookings []string

	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("There are %v tickets in total, and there are %v tickets still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("You can buy your tickets here to attend.")

	for {

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

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v, you can't book %v tickets.", remainingTickets, userTickets)
			break
		}

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v! You have successfully bought %v tickets, you should get a confirmation shortly in your email: %v.\n", firstName, lastName, userTickets, email)

		fmt.Printf("There are now %v tickets left.\n", remainingTickets)

		firstNames := []string{}

		for _, booking := range bookings { // _ is a blank identifier, it is a variable we are explicitly not using.
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The people who have booked include: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0

		if noTicketsRemaining {
			fmt.Printf("Sorry, %v tickets have been sold out. See you next year!", conferenceName)
			break
		}
	}
}
