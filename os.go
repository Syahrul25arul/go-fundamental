package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

}
