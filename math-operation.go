package main

import "fmt"

func main() {
	result := 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	a += 10
	fmt.Println(a)
	a -= 10
	fmt.Println(a)
	a *= 10
	fmt.Println(a)
	a /= 10
	fmt.Println(a)

	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	negatif := -100
	positive := +100
	fmt.Println(negatif)
	fmt.Println(positive)
}
