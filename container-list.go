package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Hendrik") // push back untuk menambahkan data paling akhir
	data.PushBack("Rizal")
	data.PushBack("Array")
	data.PushFront("Babingung") // push front untuk menambahkan data paling awal

	// fmt.Println(data.Value) // error karna harus menentukan dulu data nya dari front atau back

	fmt.Println("------")

	fmt.Println(data.Front().Next().Next().Value)

	fmt.Println(data.Front().Next().Prev().Value)

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	fmt.Println("--------")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("--------")

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

}
