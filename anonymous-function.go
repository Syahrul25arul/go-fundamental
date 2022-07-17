package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {

	blacklist := func(name string) bool { // anonymous function pada variable
		return name == "anjing"
	}

	registerUser("hendrik", blacklist)
	registerUser("anjing", blacklist)
	registerUser("anjing", func(name string) bool { // anonymous function pada argument
		return name == "anjing"
	})
}
