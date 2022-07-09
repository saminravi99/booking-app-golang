package main

import "fmt"

func main() {
	var conferenceName = "Golang Conference" // can be re assigned

	const conferenceTickets = 100 //can't change this value that is cannot re assign

	var remainingTickets = conferenceTickets - 10

	fmt.Println("Welcome to our conference booking application at", conferenceName)

	fmt.Println("Get your tickets of", conferenceName, "to attend the conference")

	fmt.Println("We have", conferenceTickets, "tickets available")

	conferenceName = "Golang Conference 2019"

	fmt.Println("The conference name is:", conferenceName)

	fmt.Println("We have", remainingTickets, "tickets remaining")

	// conferenceTickets = conferenceTickets + 100	// Error: cannot assign to conferenceTickets because it is a constant

	//Using printf method of fmt package (printf = print format)
	fmt.Printf("Welcome to %v Dhaka, Bangladesh \n", conferenceName) // \n is used to print in new line because prinnt is inlined 

	fmt.Printf("We have total %v tickets remaining out of %v tickets  \n", remainingTickets, conferenceTickets)

	//conferenceName = 50 // Error: cannot use 50 (type int) as type string in assignment

	var userName string	// type string	
	var userTickets int  // type is int because we are not giving any value to it
	//ask user for userName

	userName = "John"
	userTickets = 10

	fmt.Printf("Welcome %v \n", userName)
	fmt.Printf("You have %v tickets \n", userTickets)

	//detecting type of variable
	fmt.Printf("The type of userName is %T \n", userName)
	fmt.Printf("The type of userTickets is %T \n", userTickets)
}
