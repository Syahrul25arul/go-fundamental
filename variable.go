package main

import "fmt"

func main() {
	var name string
	name = "Hendrik Array"
	fmt.Println(name)

	name = "Rizal hendrik"
	fmt.Println(name)

	var address = "APO"
	fmt.Println(address)

	var age int8 = 30 // jika tidak di tambah tipe data yaitu int8, maka secara otomatis tipe data variable age adalah int32
	fmt.Println(age)

	// year = 2020 // akan error tidak di deklarasikan dengan var
	// fmt.Println(year)

	year := 2020
	fmt.Println(year)

	year = 2022
	fmt.Println(year)

	var (
		firstName = "Hendrik"
		lastName  = "Array"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
