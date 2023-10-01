package main

import "fmt"

func main() {
	var conferenceName = "Muffin Conference"
	const conferenceTickets = 27
	var remainingTickets = 27

	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("There are %v tickets in total, and there are %v tickets still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("You can buy your tickets here to attend.")

	var userName string
	var userTickets int
}
