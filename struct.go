package main

import "fmt"

type Customer struct {
	Name, Addres string
	Age          int
	Married      bool
}

func main() {
	var hendrik Customer
	hendrik.Name = "Hendrik Array"
	hendrik.Addres = "APO"
	hendrik.Age = 35

	jamal := Customer{ // struct Literal
		Name:   "Jamal",
		Addres: "Biak",
		Age:    26, // field int didalam template struct jika tidak di deklarasi maka secara default nilai nya 0
	}

	// budi := Customer{"Budi", "Koya", 40} // cara ini tidak di rekomendasikan karna rentan error, karna jika deklarasi field struct nya di ubah atau di tambah maka akan terjadi error

	fmt.Println(jamal)

	// fmt.Println(budi)

	fmt.Println("Name :", hendrik.Name)
	fmt.Println("Address :", hendrik.Addres)
	fmt.Println("Age :", hendrik.Age)
}
