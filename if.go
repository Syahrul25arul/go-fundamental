package main

import "fmt"

func main() {
	name := "Hendrik"

	if name == "Hendrik" {
		fmt.Println("Hello ", name)
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Saya hendrik")
	}

	if value := 5; value == 5 { // deklarasi statment di if condisition
		fmt.Println("anda benar")
	}
}
