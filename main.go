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

	var userName string
	var ticketsToBuy int

	fmt.Print("Your username: ")
	fmt.Scan(&userName) // & <-- Get memory address, pointer
	fmt.Print("Number of tickets: ")
	fmt.Scan(&ticketsToBuy)
	fmt.Println("Your username was saved in", &userName, "memory address")
	fmt.Fprintf(os.Stdout, "%v wants to buy %v tickets\n", userName, ticketsToBuy)

}
