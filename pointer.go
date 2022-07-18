package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1) // data address1 tidak berubah, padahal terjadi perubahan data di address2 yang merefrence ke address1
	/*
	* hal itu terjadi karena di golang menganut pass by value, artinya data dari address1 di duplikat lalu dimasukkan kedalam address2
	* hal itu menyebabkan address2 tidak memiliki hubungan dengan address1, karna secara default datanya di duplikat bukan refrence
	 */
	fmt.Println(address2)

	fmt.Println("---------")

	var address3 *Address = &address1 // tanda & didepan sebuan variable adalah pointer
	/*
	* pointer di gunakan untuk merefrence sebuah data dari sebuah variabel
	* data yang kita refrence adalah data yang sama dari variabel yang di refrence di lokasi memori
	* dengan menggunakan pointer maka data tidak di pass by value atau di duplikat, melainkan di refrence
	* dan jika kita mengubah data dari sebuah variable yang merefrence ke variabel lain, maka data dari variable sumber refrence nya juga akan berubah
	* hal itu disebabkan karna mereka mengakses data dari lokasi memori yang sama
	 */
	address3.City = "Papua"
	fmt.Println(address1)
	fmt.Println(address3)

	fmt.Println("---------")

	// address3 = &Address{"Jayapura", "Papua", "Indonesia"} // variable yang menjadi pointer bisa merfrence langsung ke sumber refrence nya
	/*
	* contoh nya adalah variable address3, karna sudah menjadi pointer untuk variable address1 maka variable address3 bisa di assign ulang merefrence langsung ke struct Address
	* hal ini menjadikan assign data baru yang di buat di lokasikan pada memori yang berbeda, sehingga variable address1 tidak berubah karna berbeda lokasi memori
	 */

	*address3 = Address{"Jayapura", "Papua", "Indonesia"}
	/*
	* jika menggunakan tanda * pada sebuah variable pointer, dan mengasign ulang data dari refrence
	* maka semua variable yang mengacu pada data yang sama akan pindah dan mengacu pada data yang baru di buat
	 */
	fmt.Println(address1)
	fmt.Println(address3)

	fmt.Println("--------")

	var address4 *Address = new(Address) // bisa menggunakan pointer menggunakan syntax new
	address4.City = "Wamena"
	/*
	* bedanya jika menggunakan function new maka data pointer kosong atau kembali keawal
	 */
	fmt.Println(address4)
	fmt.Println(address1)
	fmt.Println(address3)
}
