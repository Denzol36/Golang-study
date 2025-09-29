package main

import (
	"fmt"
)

func main()  {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v still avaliable\n", conferenceTickets, remainingTickets)

	var userName string 
	var userTickets uint
	var userSecondName string
	var userEmail string

	fmt.Print("Enter user first name: ")
	fmt.Scan(&userName)
	fmt.Print("Enter user second name: ")
	fmt.Scan(&userSecondName)
	fmt.Print("Enter user email name: ")
	fmt.Scan(&userEmail)
	fmt.Print("Enter amount of buying tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("\n*************\n\nUser name: %v\nUser second name: %v\nUser email: %v \nUser tickets %v\n",userName,userSecondName,userEmail,userTickets)
}