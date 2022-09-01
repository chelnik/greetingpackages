// this is my doc
package boba

import "fmt"

type Person struct {
	Name        string
	Age         int
	HomeAddress Address
}
type Address struct {
	Town         string
	street       string
	NumberOfTown int
}

// boba doc paramenrs and comments
func Boba() {
	fmt.Println("aaa boba! ")
}

func Biba() {
	fmt.Println("aaa biba! ")
}
