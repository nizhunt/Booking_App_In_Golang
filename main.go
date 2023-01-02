package main

import (
	"fmt" /* has some basic formatting functions like print, etc */
	"sync"
	"time"
)

const conferenceTickets uint = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([] UserData, 0) /* initializing an empty slice (dynamic array) of bookings where each element of the array is a mapping that contains userData */
type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main(){
	greetUsers();

	for{
		firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicket := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicket {
			bookTicket( userTickets , firstName , lastName , email)
			
			wg.Add(1)	/* Add the next one line to waitGroup */
			/* Concurrency: Create a new thread of flow for this function */
			go sendTicket( userTickets , firstName , lastName , email )

			firstNames := printFirstName()
			fmt.Printf("%v People have booked tickets. \nThese are all the people who have done the bookings: %v\n", len(firstNames), firstNames) 

			if remainingTickets == 0 {
				fmt.Printf("Sale Over! Please comeback next time.")
				break
			}
		} else{
			fmt.Printf("Input data invalid\n")
			if !isValidName {
				fmt.Printf("First name or last name too short\n")
			}
			if !isValidEmail {
				fmt.Printf("email address doesn't contain @ sign\n")
			}
			if !isValidTicket{
				fmt.Printf("Quantity of tickets is not valid\n")
			}
			continue
		}
	}
	wg.Wait()
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v of %v tickets left\n", remainingTickets,conferenceTickets)

}

func printFirstName() [] string{
	firstNames := []string{}
		/* iterate through bookings and store index and each entity called booking */

		/* "_" is a blank identifier which can be used to receive value which is not useful without generating the variable unused error */
		for _, booking := range bookings{
			firstNames = append(firstNames, booking.firstName)
		}
		return firstNames
}

func getUserInput()(string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	/* & is user to use the pointer of a variable instead of using the variable itself */
	fmt.Println("Enter Your First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your LastName")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets
	
	/* initializing an empty mapping */
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings,userData)

	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("Thankyou %v %v for booking %v tickets, you will receive a confirmation email at %v\n", firstName,lastName,  userTickets, email)
	fmt.Printf("%v tickets are remaining\n", remainingTickets) 
}

func sendTicket( userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("##########")
	fmt.Printf("Sending ticket: \n %v \n to email address %v\n", ticket, email )
	fmt.Printf("##########")

	wg.Done()

}
