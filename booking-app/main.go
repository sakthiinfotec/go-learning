package main

import (
	"booking-app/common"
	"booking-app/helper"
	"fmt"
)

// Package level variables
var conferenceName = "Go Conference"

const totalTickets uint = 50

var remainingTickets = totalTickets

var bookings = make([]common.UserData, 4)

func main() {

	greetUser()

	for {
		userName, email, userTickets := getUserInput()

		// Validate details
		isValidName, isValidEmail, isValidTickets := helper.Validate(remainingTickets, userName, email, userTickets)

		if isValidName && isValidEmail && isValidTickets {
			_, bookings = bookTicket(userName, email, userTickets)
			if remainingTickets == 0 {
				fmt.Printf("\nOur conference is booked out. Come back next year\n\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Invalid name %v. Try again with valid name", userName)
			}
			if !isValidEmail {
				fmt.Printf("Invalid email %v. Try again with valid email", email)
			}
			if !isValidTickets {
				fmt.Printf("Invalid ticket number: %v. Try again with valid ticket numbers", userTickets)
			}
		}
	}

	// Print names
	names := helper.GetUserNames(&bookings)
	fmt.Printf("Names of the bookings are: %v", names)
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

func getUserInput() (string, string, uint) {
	var userName string
	var email string
	var userTickets uint

	fmt.Print("\nEnter your name: ")
	fmt.Scan(&userName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter no of tickets: ")
	fmt.Scan(&userTickets)
	return userName, email, userTickets
}

func bookTicket(userName string, email string, userTickets uint) (uint, []common.UserData) {
	remainingTickets = remainingTickets - userTickets

	var userData = common.UserData{
		Name:       userName,
		Email:      email,
		NofTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("\nThank you %v for booking %v tickets. You will get a confirmation e-mail at %v\n", userName, userTickets, email)
	fmt.Printf("%v tickets remaining for conference %v\n", remainingTickets, conferenceName)
	return remainingTickets, bookings
}
