package main

import (
	"fmt"
	"strings"
)

const conference_tickets = 50

var conference_name = "Go Conference"
var remaining_tickets uint = 50

// making an array
var bookings = []string{}

func main() {

	greet_users()

	for true {

		first_name, last_name, email, user_tickets := get_user_input()
		// validate user iput function
		is_valid_name, is_valid_email, is_valid_ticket_num := validate_user_in(first_name, last_name, email, user_tickets, remaining_tickets)

		if is_valid_ticket_num && is_valid_email && is_valid_name {

			book_ticket(remaining_tickets, user_tickets, bookings, first_name, last_name, email, conference_name)

			// call function print first names
			var first_names = get_first_names(bookings)
			fmt.Printf("The first names of bookings are %v\n", first_names)

			if remaining_tickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !is_valid_name {
				fmt.Println("First name or last name is too short")
			}
			if !is_valid_email {
				fmt.Println("Email address does not containt '@' sign")
			}
			if !is_valid_ticket_num {
				fmt.Println("Number of tickets enter is invalid")
			}
		}

	}

}

func greet_users() {

	fmt.Printf("Welcome to %v booking application!\n", conference_name)
	fmt.Printf("We have total of: %v tickets and %v are still available\n", conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here")

}

func get_first_names(bookings []string) []string {
	first_names := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		first_names = append(first_names, names[0])
	}
	return first_names
}

func validate_user_in(first_name string, last_name string, email string, user_tickets uint, remaining_tickets uint) (bool, bool, bool) {
	is_valid_name := len(first_name) >= 2 && len(last_name) >= 2
	is_valid_email := strings.Contains(email, "@")
	is_valid_ticket_num := user_tickets > 0 && user_tickets <= remaining_tickets
	return is_valid_name, is_valid_email, is_valid_ticket_num
}

func get_user_input() (string, string, string, uint) {
	var first_name string
	var last_name string
	var email string
	var user_tickets uint

	// ask user for their name
	fmt.Println("\nEnter your first name: ") // this is a comment 'case'
	fmt.Scan(&first_name)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&last_name)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&user_tickets)

	return first_name, last_name, email, user_tickets

}

func book_ticket(remaining_tickets uint, user_tickets uint, bookings []string, first_name string, last_name string, email string, conference_name string) {
	remaining_tickets = remaining_tickets - user_tickets
	// list and arrays
	bookings = append(bookings, first_name+" "+last_name)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will revieve a comfirmation email at %v\n", first_name, last_name, user_tickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remaining_tickets, conference_name)
}
