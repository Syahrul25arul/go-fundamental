package main

import "fmt"

func main() {
	counter := 0

	increment := func() { // closure adalah sebuah  pengaksesan sebuah dari data dari luar kedalam dalam atau child nya
		// block scope adalah batasan dari ruang lingkup kerja, batasan bisa di tandai dengan adanya tanda kurung kurawal seperti function atau block if
		// dalam closure kita bisa mengakses data dari luar ke dalam selama scope di dalam adalah masih satu bagian dengan data yang di luar seperti function increment dan variable counter yang berada pada scope function main. namun kita tidak bisa mengakses data yang berada di dalam keluar
		// counter := 1 // kita bisa mendklarasikan sebuah variable di dalam scop child dengan nama variable yang sama diluar scope child. karna block dari dua scope itu berbeda seperti variable counter yang berada di scope function increment
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
