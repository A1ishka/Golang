package main

import (
	"booking/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)

var WG = sync.WaitGroup{}

func main() {
	var conferenceName = "Go Conference"
	fmt.Printf("Hello! Welcom to %v \n", conferenceName) //%T - type

	const conferenceTickets int = 50
	var remainimgTickets uint = 50
	fmt.Println("Get your tickets! We have only ", conferenceTickets, " and ", remainimgTickets, " is avalible")

	bookings := []string{} //dinamic array - string slice
	var firstName string
	var secoundName string
	var email string
	var userTickets uint
	var isInvalidName bool = true
	var isValidEmail bool = false

	/*var mymap = make(map[string]string)
	mymap["firstName"] = "fName"

	a := 5
	b := 10
	a, b = b, a*/

	for remainimgTickets > 0 || endApp == false {
		for isInvalidName == true {
			fmt.Println("\nEnter your first name: ")
			fmt.Scan(&firstName)
			fmt.Println("Enter your secound name: ")
			fmt.Scan(&secoundName)
			isInvalidName = len(firstName) < 2 || len(secoundName) < 2
		}
		for isValidEmail == false {
			fmt.Println("Enter your email: ")
			fmt.Scan(&email)
			isValidEmail = strings.Contains(email, "@")
		}
		fmt.Println("How many tickets you want to buy?")
		fmt.Scan(&userTickets)

		if userTickets > remainimgTickets {
			fmt.Printf("Cannot give you those tickets. We have only %v tickets\n", remainimgTickets)
			continue
		}

		userTickets, remainimgTickets, bookings, firstName, secoundName = helper.TicketsWork(userTickets, remainimgTickets, bookings, firstName, secoundName)
		WG.Add(1)
		go sendTickets(firstName, userTickets)
		var noTicketsRemainig bool = remainimgTickets == 0
		if noTicketsRemainig {
			fmt.Println("All were booked up. See our next time!")
			break
		}

		endingSession := "anwser"
		println("Wanna end the session?  Y/N")
		fmt.Scan(&endingSession)
		if endingSession == "Y" {
			isValidEmail = false
			isInvalidName = true
		}
	}
	WG.Wait()
}

func sendTickets(firstName string, userTickets uint) {
	time.Sleep(15 * time.Second)
	fmt.Printf("\n\n%v tickets are sending to %v\n\n", userTickets, firstName)
	WG.Done()
}
