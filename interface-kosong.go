package main

import "fmt"

type Apapun interface {
}

/*
* interface kosong adalah sebuah kontrak yang ada didalam nya
* jadi ketika kita mengimplementasikan inteface kosong maka kita bisa memanipulasi data berupa parameter, return value ataupun yang lainnya sesuka kita.
* contoh, di golang jika kita membuat function yang menerima parameter, maka kita harus secara explisit memberikan tipe data dari parameter tersebut.
* jika kita tidak memberikan tipe data pada parameter maka akan terjadi error.
* namun pada built-in function Println dari package fmt, kita bisa memasukan berbagai macam tipe data seperti string, int, struct dan lainnya.
* secara default seharusnya akan terjadi error jika pada built-in function Println dari package fmt secara explisit memberikan tipe data pada parameternya
* namun pada implementasinya, built-in function Println dari package fmt memberikan tipe data interface kosong sebagai parameternya
* hal itu membuat function Println bisa menerima argument dari tipe data yang berbeda beda
 */

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func sayGoodBy(name Apapun) {
	fmt.Println("Gooby", name)
}

func main() {
	// var data int = Ups(2) // karna type return value nya adalah interface kosong, maka saat assigment ke sebuah variabel kita tidak bisa secara explisit menentukan tipe datanya karna return valuenya bisa berbeda beda tipe data
	// var data interface{} = Ups(2) // kita bisa menggunakan tipe data interface{} seperti ini
	// var data Apapun = Ups(3)
	var data = Ups(2) // atau tidak perlu mendklarasikan tipe datanya
	fmt.Println(data)

	/*
	* say goodby menerima parameter dengan tipe data interface Apapun yang mana interface Apapun adalah interface kosong
	* secara default jika argument yang di berikan kepada function sayGoodBy bukan tipe dari implementasi interface Apapun maka akan terjadi error
	* namun karna interface Apapun adalah interface kosong, maka setiap variabel atau pun data yang telah di deklarasi akan secara otomatis mengimplementasikan interface Apapun
	* sehingga argument apapun yang akan diberikan kepada sayGoodBy akan di terima oleh function sayGoodBy
	 */
	sayGoodBy(data)

}
