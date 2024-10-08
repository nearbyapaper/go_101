package demo

import "fmt"

type str string

// func (text string) log() { // cannot define new method on non-local
// 	fmt.Println(text)
// }

func (text str) Log() {
	fmt.Println(text)
}
