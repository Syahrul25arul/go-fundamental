package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello ", firstName, " ", lastName)
}

func main() {
	/* func helloWorld() {
		fmt.Println("hello world")
	} */ // function tidak boleh di deklarasikan di dalam function main
	sayHelloTo("Hendrik", "Rizal")

}
