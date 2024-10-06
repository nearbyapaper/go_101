package demo

import (
	"fmt"
	"time"
)

type user struct {
	firstName  string
	lastName   string
	email      string
	age        int
	birthDay   string
	createDate time.Time
}

func Struct101() {
	// Create a new user
	user1 := user{
		firstName:  "John",
		lastName:   "Doe",
		email:      "johndoe@example.com",
		age:        30,
		birthDay:   "1988-05-15",
		createDate: time.Now(),
	}

	fmt.Println("================================This is Struct101================================")
	// Print the user details
	printUserDetails(user1)

	fmt.Println("\nUpdated user")
	// Update the user's email
	user1.email = "johndoe@new-example.com"
	user1.createDate = time.Now()

	// Print the updated user details
	printUserDetails(user1)
}

func printUserDetails(data user) {
	fmt.Printf("\n%s %s %d years old\n birthday %s\n email : %s\n this information update when %s", data.firstName, data.lastName, data.age, data.birthDay, data.email, data.createDate)
}
