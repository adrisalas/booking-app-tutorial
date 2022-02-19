package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Adri"
	const TOTAL_TICKETS uint = 50
	var tickets = TOTAL_TICKETS

	fmt.Print("> App initialized!\n")
	fmt.Printf("> This is a tutorial for %v's Go booking app\n\n", name)
	fmt.Printf("We have %v tickets and %v tickets are tickets are available to purchase\n", TOTAL_TICKETS, tickets)
	fmt.Printf("####################################\n\n")

	var bookings = []string{}
	var userNames []string
	for {
		var userName string
		var email string
		var ticketsToBuy int

		fmt.Print("Your username: ")
		fmt.Scan(&userName) // & <-- Get memory address, pointer
		fmt.Print("Your email: ")
		fmt.Scan(&email)
		fmt.Print("Number of tickets: ")
		fmt.Scan(&ticketsToBuy)

		tickets = tickets - uint(ticketsToBuy)
		bookings = append(bookings, userName+" "+email)
		fmt.Printf("\n%v [%v] bought %v tickets\n", userName, email, ticketsToBuy)
		fmt.Printf("There are %v tickets left\n", tickets)

		for _, booking := range bookings {
			var user = strings.Fields(booking)
			username := user[0]
			userNames = append(userNames, username)
		}
		fmt.Printf("Users that booked a ticket are: %v\n\n", userNames)
	}
}
