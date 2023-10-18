package main

import (
	"booking_app/helper"
	"fmt"

	// "strconv"
	// "strings"
	// "sync"
	"time"
)

// package level variables
// this alternative syntax cannot be used to declare constants ==> variable := value
// conferenceName := "Muffin Conference"

const conferenceTickets = 27

var conferenceName = "Muffin Conference"
var remainingTickets uint = 27

// an empty list of maps for our user data
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

// using wait-group to synchronize all our threads
// var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {

			bookTicket(userTickets, firstName, lastName, email)

			// wg.Add(1)
			// using the go keyword to create another thread of our app
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The people who have booked include: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Printf("Sorry, %v tickets have been sold out. See you next year!\n", conferenceName)
				// break
			}
		} else {
			// fmt.Printf("We only have %v, you can't book %v tickets.\nTry again please.\n", remainingTickets, userTickets)
			fmt.Println("Your input is invalid. Please Try again")

			if !isValidName {
				fmt.Println("First name and Last name must contain at least 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("Email must contain 'example@domain.com' format.")
			}
			if !isValidUserTickets {
				fmt.Println("You can only book tickets less than or equal to the available tickets.")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the %v.\n", conferenceName)
	fmt.Printf("There are %v tickets in total, and there are %v tickets still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("You can buy your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings { // _ is a blank identifier, it is a variable we are explicitly not using.
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[0]) // pick first name from list
		// firstNames = append(firstNames, booking["firstName"]) //pick first name from map
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {

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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// variableName := map[keyDataType]valueDataType
	// empty map ==> variableName := make(map[keyDataType]valueDataType)
	// // userData := make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//using struct to create a map for the user data
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings structure is %v.\n", bookings)

	fmt.Printf("Thank you %v %v! You have successfully bought %v tickets, you should get a confirmation shortly in your email: %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are now %v tickets left.\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v.\n", userTickets, firstName, lastName)
	fmt.Println("##########################################")
	fmt.Printf("Sending ticket: %v \nto email address: %v\n", ticket, email)
	fmt.Println("##########################################")
	// wg.Done()
}
