package utility

import "fmt"

func GetUserIntInput(text string) int {
	fmt.Println(text)
	var inputInt int
	fmt.Scanln(&inputInt)
	return inputInt
}

func GetUserFloatInput(text string) float64 {
	fmt.Println(text)
	var inputFloat float64
	fmt.Scanln(&inputFloat)
	return inputFloat
}

func GetUserStringInput(text string) string {
	fmt.Println(text)
	var inputString string
	fmt.Scanln(&inputString)
	return inputString
}
