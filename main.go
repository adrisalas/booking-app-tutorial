package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Adri"
	const TOTAL_TICKETS uint = 50
	var ticketsLeft = TOTAL_TICKETS

	fmt.Print("> App initialized!\n")
	fmt.Printf("> This is a tutorial for %v's Go booking app\n\n", name)
	fmt.Printf("We have %v tickets and %v tickets are tickets are available to purchase\n", TOTAL_TICKETS, ticketsLeft)
	fmt.Printf("####################################\n\n")

	var bookings = []string{}
	var userNames []string
	for ticketsLeft > 0 && len(bookings) < 50 {
		var userName string
		var email string
		var ticketsToBuy int

		fmt.Print("Your username: ")
		fmt.Scan(&userName) // & <-- Get memory address, pointer
		fmt.Print("Your email: ")
		fmt.Scan(&email)
		fmt.Print("Number of tickets: ")
		fmt.Scan(&ticketsToBuy)

		if ticketsLeft >= uint(ticketsToBuy) {
			ticketsLeft = ticketsLeft - uint(ticketsToBuy)
			bookings = append(bookings, userName+" "+email)
			fmt.Printf("\n%v [%v] bought %v tickets\n", userName, email, ticketsToBuy)
			fmt.Printf("There are %v tickets left\n\n", ticketsLeft)
		} else if ticketsLeft < uint(ticketsToBuy) {
			fmt.Println("You cannot buy that many tickets there are only", ticketsLeft, "tickets left")
		}
	}

	for _, booking := range bookings {
		var user = strings.Fields(booking)
		username := user[0]
		userNames = append(userNames, username)
	}
	fmt.Printf("Users that booked a ticket are: %v\n\n", userNames)
}
