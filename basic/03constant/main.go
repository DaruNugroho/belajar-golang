package main

import "fmt"

func main() {
	const firstName string = "Daru "
	const lastName = "Nugroho"

	// Di Golang variabel constant jika tidak dipakai maka tidak menimbulkan error
	const value = 1000

	// firstName = " Nugroho" // ERROR

	fmt.Println(firstName + lastName)
}
