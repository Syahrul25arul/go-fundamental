package main

import "fmt"

func main() {
	type noKTP string // noKTP adalah alias dari tipe data string yang
	type Married bool

	var noKtpHendrik noKTP = "2235jkhkjas09235"
	var marriedStatus bool = true
	fmt.Println(noKtpHendrik)
	fmt.Println(marriedStatus)
}
