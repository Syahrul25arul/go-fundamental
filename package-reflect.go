package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" && reflect.ValueOf(data).Field(i).Interface() == "" {
			return false
		}

		// lengthName := reflect.ValueOf(data).Field(i).Interface().(string) // .(string) di akhir digunakan untuk konversi tipe data pada tipe data interface kosong
		// fmt.Println()

		lengthString, _ := strconv.Atoi(field.Tag.Get("max"))

		// if disini untuk memvalidasi apakah panjang lebih dari tag max
		if int(len(reflect.ValueOf(data).Field(i).Interface().(string))) > lengthString {
			return false
		}
	}

	return true
}

type Sample struct {
	Name string `required:"true" max:"10"` // saat membuat tag, key tag dan : dan value jangan di pisah, contoh required:true
}

func main() {
	sample := Sample{"Hendrik Array"}

	sampleType := reflect.TypeOf(sample)
	nameField, _ := sampleType.FieldByName("Name")

	fmt.Println(string(nameField.Tag.Get("required")))
	fmt.Println(string(nameField.Tag.Get("max")))

	// sample.Name = "" // jika name nya string kosong, maka kena validasi dari is invalid dimana datanya di ambil dari reflection
	fmt.Println(isValid(sample))
}
