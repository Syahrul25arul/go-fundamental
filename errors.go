package main

import (
	"errors"
	"fmt"
)

/*
* sudah biasa di golang jika ada function yang bisa mengembalikan error maka return value nya ada dua
* yang pertama value aslinya, yang kedua errornya
 */
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan Nol ")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var contohError error = errors.New("Ups Error") // interface error sudah ada, tinggal digunakan
	fmt.Println(contohError)

	hasil, err := Pembagian(10, 100)

	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

}
