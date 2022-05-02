package helper

import(
	"fmt"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, 
	userEmail string, userTickets uint, remainingTickets uint) bool{
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")
		isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets
		if isValidName && isValidEmail && isValidTicketCount{
		return true
	} else{
		if !isValidName{
		fmt.Printf("First Name and Last Name are too short! \n")
	}
	if !isValidEmail{
		fmt.Printf("Email entered is incorrect! \n")
	}
	if !isValidTicketCount{
		fmt.Printf("Ticket count is incorrect! \n")
	}
		return false
	}
}