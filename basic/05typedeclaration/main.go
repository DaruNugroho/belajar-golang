package main

import "fmt"

func main() {
	type KTPno string
	type KTPname string
	type KTPmenikah bool

	var userKTPno KTPno = "0304489012348902"
	var userKTPname KTPname = "Daru Nugroho"
	var userKTPmarried KTPmenikah = false
	var status string = ""

	if userKTPmarried {
		status = "Menikah"
	} else {
		status = "Belum Menikah"
	}

	fmt.Println("No KTP :", userKTPno)
	fmt.Println("Nama :", userKTPname)
	fmt.Println("Menikah :", status)
}

/**
Type Declaration
- Kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
- Biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada
**/
