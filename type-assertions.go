package main

import "fmt"

func random() interface{} {
	return 32
}

func main() {
	result := random()
	var result2 = result
	result2 = "tai"

	fmt.Println(fmt.Sprintf("%T", result2))
	// fmt.Println(fmt.Sprintf("%T", result))
	// var resultString string = result.(string)b
	// fmt.Println(fmt.Sprintf("%T", resultString))
	// fmt.Println(resultString)

	fmt.Println("-----------")

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "Is string")
		fmt.Println(fmt.Sprintf("%T", value))
	case int:
		fmt.Println("Value", value, "Is string")
		fmt.Println(fmt.Sprintf("%T", value))
	default:
		fmt.Println("Unknown type")
	}
}
