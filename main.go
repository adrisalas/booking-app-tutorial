package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Adri"
	const TOTAL_TICKETS uint = 50
	var tickets = TOTAL_TICKETS

	fmt.Print("> App initialized!\n")
	fmt.Printf("> This is a tutorial for %v's Go booking app\n", name)
	fmt.Println("\nWe have", TOTAL_TICKETS, "tickets and", tickets, "tickets are available to purchase")
	fmt.Fprint(os.Stdout, "####################################\n\n")

	//var bookings [50]string ARRAY
	//var bookings []string SLICE (dynamic array)
	var bookings = []string{}

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
	bookings = append(bookings, userName+": "+email)
	fmt.Fprintf(os.Stdout, "%v bought %v tickets\n", userName, ticketsToBuy)
	fmt.Printf("The slice: %v\n", bookings)
	fmt.Printf("The slice size: %v\n", len(bookings))

}
