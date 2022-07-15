package main

import "fmt"

func main() {
	var months = [...]string{ // ... di dalam [] berfungsi untuk memberikan length dari array sesuai dengan deklarasi value dalam array. digunakan ketika bisa menentukan length dari array
		"januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fmt.Println("")
	fmt.Println("---------------")
	fmt.Println("")

	months[5] = "Diubah" // jika array sumber dari slice di ubah maka slice nya pun ikut berubah
	fmt.Println(slice1)

	slice1[0] = "Mei Lagi" // jika slice di ubah, maka array sumber juga berubah
	fmt.Println(months)
	fmt.Println(slice1)

	fmt.Println("")
	fmt.Println("---------------")
	fmt.Println("")

	var slice2 = months[8:] // jika saat slice tidak mendeklarasikan length nya, maka ketika slicenya di append. variable baru yang dibuat dari append tersebut jika ada nilainya yang di ubah, maka tidak akan mempengaruhi sumber awal array yaitu months, juga refrence slice dari varible baru yang dibuat dengan append pada slice2
	fmt.Println(slice2)

	var slice3 = append(slice2, "eko")
	fmt.Println(slice3)
	fmt.Println(slice2)

	slice3[0] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	fmt.Println("")
	fmt.Println("---------------")
	fmt.Println("")

	var slice4 = months[2:4]
	fmt.Println(slice4)

	var slice5 = append(slice4, "newMonth")
	fmt.Println(slice5)
	fmt.Println(slice4)
	fmt.Println(months)

	fmt.Println("")
	fmt.Println("---------------")
	fmt.Println("")

	slice5[1] = "newApril"
	fmt.Println(slice5)
	fmt.Println(slice4)
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(slice1)
	fmt.Println(months)

	fmt.Println("")
	fmt.Println("---------------")
	fmt.Println("")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("")
	fmt.Println("---------------")
	fmt.Println("")

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	fmt.Println("")
	fmt.Println("---------------")
	fmt.Println("")

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5} // jika lupa menambahkan ... atau length nya maka akan menjadi slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
