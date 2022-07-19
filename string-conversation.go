package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("124982133", 10, 64) // paramter kedua adalah tipe numbernya, misal oktal,decimal,binary dll, dan angka 10 maksudnya adlah parse ke bilangan decimal parametr ke tiga adalah tipe data yang akan di return

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10) // paramter kedua adalah tipe numbernya
	fmt.Println(value)
}
