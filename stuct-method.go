package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Married       bool
}

// secara default struct hanya bisa menyimpan data atau field, struct tidak bisa menyimpan function

/* func sayHello(customer Customer, name string) { // ini adalah function mandatory yang menerima struct Customer sebagai parameter
	fmt.Println("Hello", name, "My name is", customer.Name)
} */

func (customer Customer) sayHello(name string) { // ini adalah function mandatory yang dibuat seakan akan function ini milik struct Customer
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	hendrik := Customer{
		Name:    "Hendrik",
		Address: "APO",
		Age:     32,
		Married: false,
	}

	hendrik.sayHello("Dadang")
}
