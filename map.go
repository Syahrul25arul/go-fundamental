package main

import "fmt"

func main() {
	/* var person map[string]string = map[string]string {

	} */ // deklarasi panjangnya

	person := map[string]string{
		"name":    "Hendrik",
		"address": "APO",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println("")
	fmt.Println("-------------------")
	fmt.Println("")

	person["title"] = "Orang Gila"
	// person["age"] = 25 // tidak bisa mendeklarasi age dengan value int karna sudah di set tipe data string saat map pertama kali di deklarasi
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Hendrik"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "ups")
	fmt.Println(len(book))
	fmt.Println(book)
}
