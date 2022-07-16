package main

import "fmt"

func main() {
	name := "Array rs"

	switch name {
	case "Hendrik":
		fmt.Println("Hello Hendrik")
	case "Array":
		fmt.Println("Hello Array")
	default:
		fmt.Println("Hendrik Babingung")
	}

	switch length := len(name); length > 5 { // short statement dari switch, dimana bisa mendeklarasikan variable di dalam statment switch
	case true:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	} // tidak menambahkan case false karna case nya adalah boolean yang hanya mempunyai dua kondisi yaitu true atau false, jadi jika case tidak true maka secara default case nya false

	length2 := len(name)

	switch { // contonh switch tanpa kondisi
	case length2 > 10:
		fmt.Println("Nama terlalu panjang")
	case length2 > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
