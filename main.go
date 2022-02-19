package main

import (
	"fmt"
	"strings"
)

const AUTHOR = "Adri"

func main() {
	const TOTAL_TICKETS uint = 50

	greeting(TOTAL_TICKETS)

	var bookings = []string{}

	sellTickets(TOTAL_TICKETS, bookings)

	fmt.Printf("Users that booked a ticket are: %v\n\n", getUsernames(bookings))
}

func sellTickets(TOTAL_TICKETS uint, bookings []string) {
	var ticketsLeft = TOTAL_TICKETS
	for ticketsLeft > 0 && len(bookings) < 50 {
		var userName string
		var email string
		var ticketsToBuy int

		scanUserData(&userName, &email, &ticketsToBuy)

		if !isValidInput(email, ticketsToBuy) {
			continue
		}

		if ticketsLeft >= uint(ticketsToBuy) {
			ticketsLeft = ticketsLeft - uint(ticketsToBuy)
			bookings = append(bookings, userName+" "+email)
			fmt.Printf("\n%v [%v] bought %v tickets\n", userName, email, ticketsToBuy)
			fmt.Printf("There are %v tickets left\n\n", ticketsLeft)
		} else if ticketsLeft < uint(ticketsToBuy) {
			fmt.Println("You cannot buy that many tickets there are only", ticketsLeft, "tickets left")
		}
	}

}

func scanUserData(userName *string, email *string, ticketsToBuy *int) {
	fmt.Print("Your username: ")
	fmt.Scan(userName) // & <-- Get memory address, pointer
	fmt.Print("Your email: ")
	fmt.Scan(email)
	fmt.Print("Number of tickets: ")
	fmt.Scan(ticketsToBuy)
}

func greeting(tickets uint) {
	fmt.Print("> App initialized!\n")
	fmt.Printf("> This is a tutorial for %v's Go booking app\n\n", AUTHOR)
	fmt.Printf("We have %v tickets\n", tickets)
	fmt.Printf("####################################\n\n")
}

func isValidInput(email string, ticketsToBuy int) bool {
	var isValid = true
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketsToBuy := ticketsToBuy > 0
	if !isValidEmail {
		fmt.Printf("Email is not valid\n\n")
		isValid = false
	}
	if !isValidTicketsToBuy {
		fmt.Printf("Invalid number of tickets\n\n")
		isValid = false
	}
	return isValid
}

func getUsernames(bookings []string) []string {
	var userNames []string
	for _, booking := range bookings {
		var user = strings.Fields(booking)
		username := user[0]
		userNames = append(userNames, username)
	}
	return userNames
}
