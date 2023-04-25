package helper

import (
	"fmt"
	"strings"
)

func TicketsWork(userTickets uint, remainimgTickets uint, bookings []string, firstName string, secoundName string) (uint, uint, []string, string, string) {
	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, secoundName, userTickets)
	remainimgTickets -= userTickets
	fmt.Printf("\nRemaining tickets: %v", remainimgTickets)

	bookings = append(bookings, firstName+" "+secoundName)

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)

	}
	fmt.Printf("\nThere're all our bookings: %v\n", bookings)
	fmt.Printf("\nThe first names:%v\n", firstNames)

	return userTickets, remainimgTickets, bookings, firstName, secoundName
}
