package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName string = "Go conference"
var remainingTickets uint = 50
var bookings []string = []string{}

func main() {
	fmt.Println(helper.Jopa)
	greetUsers()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {

			bookTicket(userTickets, firstName, lastName, email)
			var firstNames []string = FirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			var noTicketsRemaining bool = remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our confirence is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name or last name is too short, try again")
			}
			if !isValidEmail {
				fmt.Println("Your email is incorrect, try again")
			}
			if !isValidUserTickets {
				fmt.Println("Your number of tickets is invalid, try again")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome yo %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func FirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainig for %v\n", remainingTickets, conferenceName)

}
