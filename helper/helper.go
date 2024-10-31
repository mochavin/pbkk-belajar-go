package helper

import "strings"

func ValidateName(name string) bool {
	return len(strings.TrimSpace(name)) >= 2
}

func ValidateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func ValidateTicketNumber(userTickets uint, remainingTickets uint) bool {
	return userTickets > 0 && userTickets <= remainingTickets
}
