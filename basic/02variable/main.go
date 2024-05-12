package main

import "fmt"

func main() {
	// Dekalarasi variabel biasa menggunakan 'var'
	var name string = "Daru Nugroho"
	var age int8 = 24

	// Deklarasi variabel tanpa menggunakan 'var'
	country := "Indonesia"

	fmt.Println(name)
	fmt.Println(country)

	country = "indonesia"
	fmt.Println(country)
	fmt.Println(age)
}
