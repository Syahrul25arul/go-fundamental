package other

import "fmt"

/*
* pembuktian bahwa dua package yang berbeda bisa memiliki function yang sama
 */
func SayHello(name string) {
	fmt.Println("Hello", name)
}
