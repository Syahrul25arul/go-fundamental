package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func main() { // function main adalah function yang akan di eksekusi sama golang secara default
	sayHello() // kenapa function sayHello di panggin di dalam function main, karna function main adalah tempat pertama kode yang dieksekusi golang
	sayHello()
}
