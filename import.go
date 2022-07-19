package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Hendrik")
	// helper.sayGoodBy("hendrik") // function yang di awali dengan huruf kecil, tidak bisa digunakan di package lain

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // variabel yang di awali dengan huruf besar tidak bisa digunakan di pacakge lain

}
