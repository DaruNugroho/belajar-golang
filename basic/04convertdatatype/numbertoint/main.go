package main

import "fmt"

func main() {
	var nilai32 int32 = 132
	var nilai64 int64 = int64(nilai32) // convert int32 -> int64
	var nilai8 int8 = int8(nilai32)    // terjadi integer overflow

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
}
