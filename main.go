package main

import (
	"fmt"
	"strconv"
	"github.com/Ashutosh-Chauhan93/booking-app/helper"
)

var conferenceName = "Go Conference"
var conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {
		firstName, lastName, userEmail, userTickets := enterUserInput()
		if helper.ValidateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets){
			bookTickets(firstName, lastName, userEmail, userTickets)
			firstNames := getFirstNames()
			fmt.Printf("These first names of bookings are %v\n", firstNames)
			
			var noticketsremaining bool = remainingTickets == 0
			if noticketsremaining{
				fmt.Println("Our conference is booked now!!, Please come next year")
				break
			}
		
		} else{
			continue
		} 

	}

}

func greetUsers(){
	fmt.Printf("Hello! Welcome to %v ticket booking app\n", conferenceName)
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, conferenceTickets, conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are avialable...\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!!")
}

func getFirstNames() []string{
	firstNames := []string{}

	for _, booking := range bookings{
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}


func enterUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint

	fmt.Println("Enter your firstName")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func bookTickets(firstName string, lastName string, userEmail string, userTickets uint){
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["userEmail"] = userEmail
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you User %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining\n", remainingTickets)

}