package demo

import "fmt"

func Pointers101() {
	// Declare an integer variable
	var age int = 10

	// Declare a pointer to an integer variable
	var agePtr *int = &age

	// Print the address of the integer variable
	fmt.Println("Address of age:", &age) // '&' reference to variable to the address

	// Print the value of the integer variable using the pointer
	fmt.Println("Value of age using ptr:", *agePtr) // '*' dereference pointer to the value

	if age == *agePtr {
		fmt.Println("'age' and '*agePtr' Both variables are equal")
	}

	if &age == agePtr {
		fmt.Println("'&age' and 'agePtr' Both variables are pointing to the same memory location")
	}

	fmt.Println("Adult year by A way:", getAdultYear(&age))
	fmt.Println("Adult year by B way:", getAdultYear(agePtr))
}

func getAdultYear(age *int) int {
	return *age - 18
}
