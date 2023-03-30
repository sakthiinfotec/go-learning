package helper

import (
	"booking-app/common"
	"strings"
)

func Validate(remainingTickets uint, userName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(userName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}

func GetUserNames(bookings *[]common.UserData) []string {
	names := []string{}
	for _, booking := range *bookings {
		names = append(names, booking.Name)
	}
	return names
}
