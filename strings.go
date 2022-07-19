package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hendrik array", "Hendrik"))
	fmt.Println(strings.Split("Hendrik rizal array", " "))

	fmt.Println(strings.ToLower("Hendrik Rizal Array"))
	fmt.Println(strings.ToUpper("Hendrik Rizal Array"))
	fmt.Println(strings.ToTitle("hendrik rizal array"))
	fmt.Println(strings.ToTitle("hendrik rizal array"))
	fmt.Println(strings.Trim(" hendrik rizal array  ", " "))

	fmt.Println(strings.ReplaceAll("Hendrik Rizal Rizal Array", "Rizal", "Babingung"))
}
