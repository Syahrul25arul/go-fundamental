package helper

import "fmt"

/*
* package helper menandakan bahwa source code ini sedang berada di direktori helper
* dalam sebuah package atau helper, tidak boleh ada nama function yang sama
* nama function bisa sama pada package atau direktori yang berbeda
* contoh package other dan helper yang sama sama mempunyai function Sayhello
 */

func SayHello(name string) {
	fmt.Println("Hello", name)
}
