package main

import (
	"fmt"
	"strings"
	"time"
)

const AUTHOR = "Adri"

type User struct {
	userName     string
	email        string
	ticketsToBuy int
}

func main() {
	const TOTAL_TICKETS uint = 50

	greeting(TOTAL_TICKETS)

	var bookings = []User{}

	sellTickets(TOTAL_TICKETS, &bookings)

	fmt.Printf("Users that booked a ticket are: %v\n\n", getUsernames(bookings))
}

func sellTickets(TOTAL_TICKETS uint, bookings *[]User) {
	var ticketsLeft = TOTAL_TICKETS
	for ticketsLeft > 0 && uint(len(*bookings)) < TOTAL_TICKETS {
		var user User
		scanUserData(&user)
		fmt.Println(user)

		if !isValidInput(&user) {
			continue
		}

		if ticketsLeft >= uint(user.ticketsToBuy) {
			ticketsLeft = bookTickets(ticketsLeft, user, bookings)
			go sendTickets(user) // This is a Runnable/Thread WHATTT MINDBLOWN
		} else if ticketsLeft < uint(user.ticketsToBuy) {
			fmt.Println("You cannot buy that many tickets there are only", ticketsLeft, "tickets left")
		}
	}

}

func scanUserData(userPointer *User) {
	// TODO There should be another form in Go to do this "C Trick" to work with memory
	fmt.Print("Your username: ")
	fmt.Scan(&(*userPointer).userName)
	fmt.Print("Your email: ")
	fmt.Scan(&(*userPointer).email)
	fmt.Print("Number of tickets: ")
	fmt.Scan(&(*userPointer).ticketsToBuy)
}

func greeting(tickets uint) {
	fmt.Print("> App initialized!\n")
	fmt.Printf("> This is a tutorial for %v's Go booking app\n\n", AUTHOR)
	fmt.Printf("We have %v tickets\n", tickets)
	fmt.Printf("####################################\n\n")
}

func isValidInput(userPointer *User) bool {
	var user User = *userPointer
	var isValid = true
	isValidEmail := strings.Contains(user.email, "@") && strings.Contains(user.email, ".")
	isValidTicketsToBuy := user.ticketsToBuy > 0
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

func getUsernames(bookings []User) []string {
	var userNames []string
	for _, booking := range bookings {
		userNames = append(userNames, booking.userName)
	}
	return userNames
}

func bookTickets(ticketsLeft uint, user User, bookings *[]User) uint {
	ticketsLeft = ticketsLeft - uint(user.ticketsToBuy)
	*bookings = append(*bookings, user)
	fmt.Printf("\n%v [%v] bought %v tickets\n", user.userName, user.email, user.ticketsToBuy)
	fmt.Printf("There are %v tickets left\n\n", ticketsLeft)
	return ticketsLeft
}

func sendTickets(user User) {
	//Simulate generation of ticket
	var ticket = fmt.Sprintf("%v tickets for %v", user.ticketsToBuy, user.userName)
	//Simulate sending in a email
	time.Sleep(10 * time.Second)
	fmt.Print("\n#######################\n")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, user.email)
	fmt.Print("#######################\n\n")
}
