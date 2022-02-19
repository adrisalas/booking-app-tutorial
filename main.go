package main

import "fmt"

func main() {
	var name = "Adri"
	const TOTAL_TICKETS = 50
	var tickets = TOTAL_TICKETS

	fmt.Print("App initialized! ")
	fmt.Printf("This is a tutorial for %v's Go booking app\n", name)
	fmt.Println("\nWe have", TOTAL_TICKETS, "tickets and", tickets, "tickets are available to purchase")

}
