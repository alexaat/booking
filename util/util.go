package util

import "fmt"

func IsValidNumberOfTickets(ticketsInputFromUser int, remainingTickets uint) bool {
	if ticketsInputFromUser <= 0 {
		fmt.Println(InvalidNumberOfTickets)
		return false
	}
	if ticketsInputFromUser > int(remainingTickets) {
		fmt.Printf(TooManyTickets, remainingTickets, ticketsInputFromUser)
		return false
	}
	return true
}
