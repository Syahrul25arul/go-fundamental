package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) { // tanda bintang di paramerter menunjukkan kalau tipe datanya harus pointer
	address.Country = "indonesia"
}

func main() {
	var alamat = Address{
		City:     "Jayapura",
		Province: "Papua",
		Country:  "",
	}

	ChangeAddressToIndonesia(&alamat) // tanda & membuat argument menjadi pointer, namun kalau variable nya sudah pointer tidak perlu menambahkan &
	fmt.Println(alamat)
}
