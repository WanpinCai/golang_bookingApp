package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
// var bookings = []string{}
// var bookings = make([]map[string]string,0) //list of map
var bookings = make([]UserData,0) 

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main(){

	greetUsers()

	// for {
		firstName,lastName,email,userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName,lastName,email,userTickets,remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket( userTickets,  firstName, lastName, email)
			// go sendTicket(userTickets,  firstName, lastName, email) //concurrency
			wg.Add(1)
			go sendTicket(userTickets,  firstName, lastName, email) //concurrency

			// call function print first name
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n",firstNames)

			if remainingTickets == 0{
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				// break
			}
		}else{
			if !isValidName{
				fmt.Println("Your first name or last name input data is invalid")
			}
			if !isValidEmail {
				fmt.Println("Your email input is invalid")
			}
			if !isValidTicketNumber{
				fmt.Println("number of tickets your entered is invalid")
			}
		}
		wg.Wait()
	}
// }

func greetUsers(){
	fmt.Printf("Welcome to %v our conference\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings{
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput()(string,string,string,uint){
	var firstName string
		var lastName string
		var email string
		var userTickets uint 
		//ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		return firstName,lastName,email,userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string ){
	remainingTickets = remainingTickets - userTickets
	// create a map for a user
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n",bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets,firstName,lastName)
	fmt.Println("---------------------------------------------------")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket,email)
	fmt.Println("---------------------------------------------------")
	wg.Done()
}

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets),10)

	// var myslice []string
	// var mymap map[string]string
// ➜  booking-app go run main.go helper.go
	// fmt.Printf("The first names of bookings are: %v\n",firstNames)

		// isValidCity := city == "Singapore" || city == "London"
		// !isValidCity
		// isInvalidCity := city != "Singapore" && city != "London"

			// fmt.Println("Your input data is invalid")
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n",remainingTickets,userTickets)

	// city := "London"
	// switch city {
	// case "New York":
	// 	// execute code for booking New York conference tickets
	// case "Singapore":
	// 	// execute code for booking Singapore conference tickets
	// case "Dublin":
	// 	// execute code for booking Dublin conference tickets
	// case "Beijing","Hongkong":
	// 	// execute code for booking New York conference tickets
	// case "London":
	// 	// execute code for booking New York conference tickets
	// default:
	// 	fmt.Print("No valid city selected")

	// }