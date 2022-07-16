package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	}
	return "Hello " + name
}

func main() {
	result := getHello("Hendrik")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
