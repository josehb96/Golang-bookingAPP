package main

import (
	"fmt"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	//var bookings [50]string // ARRAY
	//var bookings = [50]string{"Jose", "Carlos", "Pepe"} // ARRAY
	var bookings []string // SLICE
	//var bookings []string{} // Alternative way to declare a slice
	//bookings := []string{} // Alternative way to declare a slice

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName // Working with arrays
	bookings = append(bookings, firstName+" "+lastName) // Working with slices

	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Array type: %T\n", bookings)
	// fmt.Printf("Array length: %v\n", len(bookings))

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank your %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v\n", bookings)

}
