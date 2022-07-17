package main

import "fmt"

type HasName interface { // interface adalah sebuah tipe data abstract yang berisikan definisi dari method method
	GetName() string
}

type HasFirstName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func SayFirstName(hasFirstName HasFirstName) {
	fmt.Println("This is my firstName", hasFirstName.GetName())
}

type Person struct {
	Name string
}

/*
* di golang kita tidak bisa mengikatkan sebuah tipe data secara explisit ke sebuah interface
* namun jika kita ingin mengikatkan sebuah tipe data seperti struct ke sebuah interface maka kita hanya perlu mengimplementasikan isi dari sebuah interface
* golang secara otomatis akan membaca apakah sebuah tipe data atau struct mengimplementasikan semua method dari sebuah interface
* jika iya, maka golang secara otomatis akan mendefinisikan bahwa tipe data tersebut telah mengimplementasikan sebuah interface
* contoh pada struct Person, telah mendefinisikan funtion GetName dimana interface HasName memiliki kontrak function GetNameconst
* maka secara otomatis golang akan mendefinisikan bahwa struct Person telah mengimplementasikan interface HasName
 */
func (person Person) GetName() string {
	return person.Name
}

/*
* pada contoh di atas, bisa dilihat bahwa sebuah interface yang berbeda bisa memiliki kontrak method yang sama
* dan jika isi dari kontrak method nya sama maka jika ada sebuah struct yang mengimplementasikan method dari kontrak tersebut
* maka secara otomatis struct tersebut akan mengimplementasika beberapa interface yang memiliki kontrak method yang sama
* contoh pada struct Person yang mengimplementasika method GetName, dimana method GetName adalah isi kontrak dari interface HasName dan HasFirstName
* maka secara otomatis struct Person telah mengimplementasikan interface HasName dan HasFirstName
 */
func main() {
	var hendrik Person
	hendrik.Name = "Hendrik"
	fmt.Println(hendrik)
	SayHello(hendrik)
	SayFirstName(hendrik)
}
