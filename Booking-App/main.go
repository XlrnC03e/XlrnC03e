package main

//Remember Nana Viedo!!!!! - Time Stamp: 2:53:20
import (
	"Booking-app/helper"
	"fmt"
	"strconv"
) 


	var conferenceName = "Go Conference"
	const conferenceTickets = 50 
	var remainingTickets  uint = 50
    var bookings = make([]map[string]string, 0)


 func main(){
    
    greetUsers()
	
	for{


	firstName, lastName, email, userTickets := getUserInput()
	isValidEmail, isValidName, isValidTicketNumber := helper.ValidateInput(firstName, lastName, email, userTickets,remainingTickets)



	if isValidName && isValidEmail && isValidTicketNumber {
		
	bookTicket(userTickets, firstName, lastName, email)
	
	firstNames := getsFirstNames()
	fmt.Printf("\nThe first names of bookings are: %v\n" , firstNames)

	
	if remainingTickets == 0 {
		//end program
		fmt.Println("Our Confrenece is booked out. Come back next year.")
		break

	}

}else{
   if !isValidName {
	   fmt.Println("first name or last name you entered is too short")
   }
   if !isValidEmail{
	     fmt.Println("The email address you entered doesnt contain '@' sign ")

   }
   if !isValidTicketNumber{

	fmt.Printf("The amount of tickets you entered excceded the current amount we have, which is %v", remainingTickets)
   }
	
		continue
   }

	
	 
	
	
}

}

	func greetUsers(){
		fmt.Printf("Welcome to %v!\n", conferenceName)
		fmt.Printf("We have total of  %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
		fmt.Println("Order your tickets here!!!") 

	}

	func getsFirstNames() []string{
		firstNames := []string{}
	for _, booking := range bookings {

		
		firstNames = append(firstNames, booking["firstName"])
	}
return firstNames	
} 


func getUserInput()(string,string,string,uint){
       
		var firstName string
	
		var lastName string
	
		var email string
	
		var userTickets uint
	
	
	//ask user for their name
	fmt.Println("\nEnter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("\nEnter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("\nEnter your Email: ")
	fmt.Scan(&email)
	fmt.Println("\nEnter the number of Ticket(s) you want: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets-= uint(userTickets)

	
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	
	fmt.Printf("Thank you %v %v for ordering %v tickets! You will recieve the confirmation email at %v\n", firstName, lastName, userTickets,email)
    fmt.Printf("%v tickets remaining for %v",remainingTickets, conferenceName)

}