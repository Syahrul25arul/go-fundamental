package main

import "fmt"

func getBiodata() (string, string, int) {
	return "Hendrik", "Array", 38
}

func showBiodata(firstName, lastName string, age int) { // jika parameter function awal tidak di berikan tipe data maka dia akan mengikuti tipe data parameter setelahnya
	fmt.Println("First Name : ", firstName) // jika tidak sama sekali memberikan tipe data di semua paramter maka akan terjadi error
	fmt.Println("Last Name : ", lastName)
	fmt.Println("Age : ", age)

}

func main() {
	var firstName, lastName, age = getBiodata()
	// fmt.Println(lastName)
	// fmt.Println(firstName)
	// fmt.Println(age)

	// fmt.Println("First Name = ", firstName, ", Last Name = ", lastName, ", Age ", age)
	showBiodata(firstName, lastName, age)

	nameFirst, _, _ := getBiodata() // jika hanya ingin mengambil beberapa return value dari function maka gunakan _ untuk menandai bahwa return value tersebut tidak di butuhkan
	// namun harus minimal menangkap satu dari sekian return value yang di kembalikan, jika tidak akan terjadi erro
	fmt.Println(nameFirst)
}
