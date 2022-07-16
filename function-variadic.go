package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {

	total := sumAll(2, 4, 6, 8, 2, 3)
	// total := sumAll() // function variadic bisa menerima argument minimal 0 dan maksimal tak terbatas
	fmt.Println(total)

	slice := []int{10, 20, 30, 40, 50}
	// total = sumAll(slice) // ini akan terjadi error karna argument yang kita masukkan adalah slice sedangkan tipe data dari parameter function sumAll adalah int
	total = sumAll(slice...) // ... di dalam argument gunanya untuk memecah dan mengeluarkan values dari slice
}
