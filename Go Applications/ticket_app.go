package main

import (
	"fmt"
	"strings"
)

func main() {

	conference_name := "Go Conference"
	const conference_tickets = 50
	var remaining_tickets uint = 50
	// making an array
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conference_name)
	fmt.Printf("We have total of: %v tickets and %v are still available\n", conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here")

	for true {
		var first_name string
		var last_name string
		var email string
		var user_tickets uint

		// ask user for their name
		fmt.Println("\nEnter your first name: ")
		fmt.Scan(&first_name)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&last_name)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&user_tickets)

		is_valid_name := len(first_name) >= 2 && len(last_name) >= 2
		is_valid_email := strings.Contains(email, "@")
		is_valid_ticket_num := user_tickets > 0 && user_tickets <= remaining_tickets

		if is_valid_ticket_num && is_valid_email && is_valid_name {
			remaining_tickets = remaining_tickets - user_tickets
			// list and arrays
			bookings = append(bookings, first_name+" "+last_name)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will revieve a comfirmation email at %v\n", first_name, last_name, user_tickets, email)

			fmt.Printf("%v tickets remaining for %v\n", remaining_tickets, conference_name)

			first_names := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				first_names = append(first_names, names[0])
			}
			fmt.Printf("The first name of bookings are: %v\n", first_names)

			if remaining_tickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			fmt.Printf("Your input data is invalid, try again!")
		}

	}

}
