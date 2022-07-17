package main

import "fmt"

type Filter func(string) string // ini adalah type declaration untuk mengaliaskan function dengan parameter string dan tipe data string

func sayHelloWithFilter(name string, filter Filter) { // parameter kedua adalah parameter yang menerima function, dimana function tersebut menerima parameter string dan mengembalikan return value string
	nameFilter := filter(name)
	fmt.Println("Hello ", nameFilter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Hendrik", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
