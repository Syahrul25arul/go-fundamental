package main

import "fmt"

func getGoodBy(name string) string {
	return "Good Bye " + name
}

func main() {
	goodBy := getGoodBy // keuntungan memasukkan function ke dalam variabel adalah bisa di masukkan lagi sebagai argument kefunction lain

	fmt.Println(goodBy("Hendrik"))
}
