package main

import "fmt"

func main() {
	var name1 = "Eko"
	var name2 = "Eko"
	var name3 = "Budi"

	var result bool = name1 == name2
	var result2 bool = name1 > name3
	fmt.Println(result)
	fmt.Println(result2)

	var value1 = 100
	value2 := 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
