package main

import "fmt"

func main() {
	var names [3]string // variable array pada golang harus di deklarasikan length dan tipe datanya, jika tidak akan terjadi error

	names[0] = "Rizal"
	names[1] = "Hendrik"
	names[2] = "Array"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		95,
	} // values dari index terakhir harus di akhiri dengan koma, jika tidak akan terjadi error
	fmt.Println(values)

	values[2] = 200
	fmt.Println(values)
	fmt.Println(len(names))
	fmt.Println(len(values))

}
