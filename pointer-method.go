package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() { // method tanpa pointer
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method", man.Name) // di method berubah namun di struct di main tidak karna pass by value
}

func (man *Man) Gender() { // method dengan pointer
	man.Name = "Men " + man.Name // data struct di main akan berubah karna pass by refrence
}

func main() {
	hendrik := Man{"Hendrik"}
	hendrik.Married()
	fmt.Println(hendrik.Name)

	fmt.Println("-----")
	hendrik.Gender()
	fmt.Println(hendrik.Name)
}
