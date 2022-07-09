package main

import "fmt"

func main() {
	var conferenceName = "Golang Conference"

	const conferenceTickets = 100

	fmt.Println("Welcome to our conference booking application at", conferenceName)

	fmt.Println("Get your tickets of", conferenceName, "to attend the conference")

	fmt.Println("We have", conferenceTickets, "tickets available")

	conferenceName = "Golang Conference 2019"

	fmt.Println("The conference name is:", conferenceName)

	// conferenceTickets = conferenceTickets + 100	// Error: cannot assign to conferenceTickets because it is a constant
}
