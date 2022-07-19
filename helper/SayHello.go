package helper

import "fmt"

// var version = 1 // variabel yang di awali dengan huruf besar tidak bisa digunakan di pacakge lain
var Application = "Belajar golang"

/*
* package helper menandakan bahwa source code ini sedang berada di direktori helper
* dalam sebuah package atau helper, tidak boleh ada nama function yang sama
* nama function bisa sama pada package atau direktori yang berbeda
* contoh package other dan helper yang sama sama mempunyai function Sayhello
 */

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBy(name string) { // function yang di awali dengan huruf kecil tidak bisa digunakna di huruf lain
	fmt.Println("Gooby", name)
}
