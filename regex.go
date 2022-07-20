package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	var regexEamil *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9 \_]+@[a-z]{2,6}\.+[a-z]{2,4}$`)

	fmt.Println("hendrik@gmail.com", regexEamil.MatchString("hendrik@gmail.com"))
	fmt.Println("pas25kata@gmail.com", regexEamil.MatchString("pas25kata@gmail.com"))
	fmt.Println("syahrul25arul@gmail.com", regexEamil.MatchString("syahrul25arul@gmail.com"))
	fmt.Println("arulprojectdualima@gmail.com", regexEamil.MatchString("arulprojectdualima@gmail.com"))
	fmt.Println("hendrik.dualima@gmail.com", regexEamil.MatchString("hendrik.dualima@gmail.com"))
	fmt.Println("hendrik25dualima@g.com", regexEamil.MatchString("hendrik25dualima@g.com"))
	fmt.Println("hendrik25dualima@g.c", regexEamil.MatchString("hendrik25dualima@g.c"))
	fmt.Println("hendrik25dualima@gmail.c", regexEamil.MatchString("hendrik25dualima@gmail.c"))
	fmt.Println("hendrik25dualima@gmail.comsat", regexEamil.MatchString("hendrik25dualima@gmail.comsat"))

	// fmt.Println(regex.MatchString("ebo"))
	// fmt.Println(regex.MatchString("eto"))
	// fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("eko eka edo eto eyo eki", 2) // jika paramter kedua -1 maka akan di cari sampe string terakhir
	fmt.Println(search)
}
