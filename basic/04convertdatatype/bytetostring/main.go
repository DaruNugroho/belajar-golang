package main

import "fmt"

func main() {
	var name = "Daru Nugroho"
	var byte byte = name[0]
	var byteToChar string = string(byte)

	fmt.Println(name)
	fmt.Println(byte)
	fmt.Println(byteToChar)
}
