package main

import (
	"booking/util"
	"fmt"
	"strings"
	"sync"
	"time"
)

//var bookingsArray [50]string

var conferenceName = "Go Conference"
var invalidName = "Invlid name. Please try again"
var invalidEmail = "Invlid email. Please try again"
var soldOut = "All tickest are sold out. See you next year!"

const TOTAL_TICKETS uint = 50

var remainingTickets uint = TOTAL_TICKETS
var firstName string
var lastName string
var email string
var tickets uint
var bookings []UserData

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var wg = sync.WaitGroup{}

func main() {

	greetings()

	for remainingTickets > 0 {

		firstName = getUserName("first")
		lastName = getUserName("last")
		email = getUserEmail()
		tickets = getNumberOfTickets(remainingTickets)

		completeBooking()

		printLogs()

	}

	fmt.Println(soldOut)
	wg.Wait()
}

func greetings() {
	fmt.Printf("Welcome to %v booking system!\n", conferenceName)
}

func getUserName(firstOrLast string) string {
	var name string
	for {
		fmt.Printf("Enter your %v name: ", firstOrLast)
		fmt.Scan(&name)
		if len(name) > 1 {
			break
		} else {
			fmt.Println(invalidName)
		}
	}
	return name
}

func getUserEmail() string {
	var email string
	for {
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		if strings.Contains(email, "@") {
			break
		} else {
			fmt.Println(invalidEmail)
		}
	}
	return email
}

func getNumberOfTickets(remainingTickets uint) uint {
	var ticketsInputFromUser int
	var tickets uint

	for {
		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&ticketsInputFromUser)

		if util.IsValidNumberOfTickets(ticketsInputFromUser, remainingTickets) {
			tickets = uint(ticketsInputFromUser)
			break
		}
	}
	return tickets
}

func completeBooking() {
	remainingTickets = remainingTickets - tickets
	var user = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   tickets,
	}
	var fullName = firstName + " " + lastName
	bookings = append(bookings, user)
	fmt.Printf("%v booked %v tickets. Confirmation is send to %v\n", fullName, tickets, email)

	wg.Add(1)
	go generateAndSendTicket(user)

}

func generateAndSendTicket(user UserData) {
	time.Sleep(45 * time.Second)
	fmt.Printf("\nSending ticket to: %v\n", user.email)
	fmt.Println("###################")
	fmt.Println(conferenceName)
	var ticket = fmt.Sprintf("Booking for %v %v\nNumber of places %v ", user.firstName, user.lastName, user.tickets)
	fmt.Println(ticket)
	fmt.Println("###################")
	wg.Done()
}

func printLogs() {
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	fmt.Printf("Bookings size: %v\n", len(bookings))
	fmt.Printf("All bookings: %v\n", bookings)
}
