package demo

import (
	"fmt"
	"time"
)

type user struct { // struct blueprints
	firstName  string
	lastName   string
	email      string
	age        int
	birthDay   string
	createDate time.Time
}

func (data user) printUserDetails() {
	// (data user) calles 'receiver argument'
	// which is not specified here in regular parameter list
	// but which(receiver argument) instead is specified here
	// in frontof the functionname to, as mentioned above
	fmt.Printf("\n%s %s %d years old\n birthday %s\n email : %s\n this information update when %s\n", data.firstName, data.lastName, data.age, data.birthDay, data.email, data.createDate)
}
func (data *user) printUserDetailsByPointer() {
	// in GO Struct have improved you do not have to dereference for access data pointer in struct
	fmt.Println("====================================================")
	fmt.Printf("\n%s %s %d years old\n birthday %s\n email : %s\n this information update when %s\n", data.firstName, data.lastName, data.age, data.birthDay, data.email, data.createDate)
	// but in real process you have to do like this
	// fmt.Println("====================================================")
	// fmt.Printf("\n%s %s %d years old\n birthday %s\n email : %s\n this information update when %s\n", (*data).firstName, (*data).lastName, (*data).age, (*data).birthDay, (*data).email, (*data).createDate)
}

func (data user) clearUserName() { // use case It's not work for clear struct
	data.firstName = ""
	data.lastName = ""
}

func (data *user) clearUserNameByPointer() { // use case It's work for clear struct by using pointer
	// (data *user) let this method edit the original Struct
	data.firstName = ""
	data.lastName = ""
}

func newUserPoorWay(firstName string, lastName string, email string, age int, birthDay string) user {
	return user{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		age:        age,
		birthDay:   birthDay,
		createDate: time.Now(),
	}
}

func newUserGoodWay(firstName string, lastName string, email string, age int, birthDay string) *user {
	return &user{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		age:        age,
		birthDay:   birthDay,
		createDate: time.Now(),
	}
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
	} // Literal Notation

	fmt.Println("================================This is Struct101================================")
	// Print the user details
	// printUserDetails(user1)
	user1.printUserDetails()

	fmt.Println("\nUpdated user")
	// Update the user's email
	user1.email = "johndoe@new-example.com"
	user1.createDate = time.Now()

	// Print the updated user details
	fmt.Println("================================Use StructPointer Pass to Function================================")
	// printUserDetailsByPointer(&user1)
	// user1.clearUserName() // not working cause not edit real struct
	fmt.Println("Clear User Name")
	user1.clearUserNameByPointer()

	user1.printUserDetailsByPointer()

	// Poor way to use contructor function
	user2 := newUserPoorWay("Mohamed", "Salah", "mo_salah@gmail.com", 32, "June 15, 1992")
	user2.printUserDetails()

	// Good way to use contructor function
	// var user3 *user
	user3 := newUserGoodWay("Luis", "Diaz", "lucho@gmail.com", 27, "January 13, 1997")
	user3.printUserDetails()
}

// func printUserDetailsByPointer(data *user) {
// 	// in GO Struct have improved you do not have to dereference for access data pointer in struct
// 	fmt.Println("====================================================")
// 	fmt.Printf("\n%s %s %d years old\n birthday %s\n email : %s\n this information update when %s\n", data.firstName, data.lastName, data.age, data.birthDay, data.email, data.createDate)
// 	// but in real process you have to do like this
// 	fmt.Println("====================================================")
// 	fmt.Printf("\n%s %s %d years old\n birthday %s\n email : %s\n this information update when %s\n", (*data).firstName, (*data).lastName, (*data).age, (*data).birthDay, (*data).email, (*data).createDate)
// }
