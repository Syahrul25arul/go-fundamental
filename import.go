package main

import (
	_ "belajar-golang-dasar/database" // tanda _ digunakan untuk membiarkan package yang tidak kita tetap pake tetap ada
	"belajar-golang-dasar/helper"     // karna secara default golang akan komplain jika ada package yang di import namun tidak digunakan
	"fmt"
)

func main() {
	helper.SayHello("Hendrik")
	// helper.sayGoodBy("hendrik") // function yang di awali dengan huruf kecil, tidak bisa digunakan di package lain

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // variabel yang di awali dengan huruf besar tidak bisa digunakan di pacakge lain

	// result := database.GetDatabase() // func init di package database akan di eksekusi pertamakali secara otomatis ketika di akses
	// fmt.Println(result)
}
