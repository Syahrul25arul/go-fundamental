package main

import "fmt"

func main() {
	const firstName string = "Hendrik"
	const lastName = "Array"
	const value = 1000

	// variable constant tidak error jika tidak digunakan, berbeda dengan variable biasa

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		addres string = "APO"
		city          = "Jayapura"
	)

	fmt.Println(addres)
	fmt.Println(city)
}
