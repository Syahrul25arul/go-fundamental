package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // tipe data int8 hanya mampu menampung angka sampai 100, akibat nya data akan di ulang dari awal range tipe datanya
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) // hasil -96

	var name = "Eko"
	var e byte = name[0] // variable e menampung byte dari name[0] atau karakter "E", sehingga variable e harus di konversi terlebih dahulu untuk menampilkan string nya
	var eString = string(e)

	fmt.Println(eString)
}
