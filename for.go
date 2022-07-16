package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 10 { // perulangan akan berakhir ketika kondisi bernilai false
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("-----------")

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke ", i)
	}

	fmt.Println("-----------")

	slice := []string{
		"hendrik",
		"rizal",
		"array",
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("--------")

	for index, value := range slice { // for range untuk iterasi slice
		fmt.Println("Index ", index, " = ", value)
	}

	fmt.Println("--------")

	for _, value := range slice { // karna di golang variable yang telah di deklarasi jika tidak digunakan akan terjadi error, maka untuk mengatasi itu di gunakan _ dalam perulangan for ranges
		fmt.Println(value)
	}

	fmt.Println("--------")

	person := map[string]string{
		"name":    "hendrik",
		"address": "APO",
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

	fmt.Println("--------")

	for _, value := range person {
		fmt.Println(value)
	}
}
