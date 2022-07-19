package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	var username *string = flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	number := flag.Int("number", 100, "Put your number")

	/*
	* flag.string atau flag.parse digunakan untuk menagkap data path dari command line
	* flag parse di gunakan untuk memparsing data yang telah ditangkap dari command line
	 */

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *username)
	fmt.Println("password : ", *password)
	fmt.Println("number : ", *number)
}
