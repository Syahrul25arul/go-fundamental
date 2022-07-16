package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Rizal"
	middleName = "Hendrik"
	lastName = "Array"

	// return firstName, middleName, lastName
	return // jika memberikan nama di tipe return value maka kita tidak perlu mendeklarasikan variable yang akan di return karna sudah terdaftar di variable return value
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	// a, b, c := getCompleteName()  // nama variable yang di assigment tidak harus mengikuti nama variable yang ada di return value

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}
