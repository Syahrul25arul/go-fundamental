package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() // defer digunakan untuk memaksa sebuah function untuk tetap dijalankan di akhir statement, walaupun saat pertengahan eksekusi terjadi error, dengan defer function tersebut akan tetap di jalanakna di akhir statment
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("result", result)

}

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover() // recover di gunakan untuk menangkap data panic dan membatalkan pemberhentian eksekusi code dari bawaan function panic. kenapa recover function tidak di letakkan pada function runApp di bawah function panic, karena saat function panic di eksekusi maka secara otomatis akan memberhenti paksakan program yang menyebabkan script di bawahnya tidak di eksekusi sehingga function recover tidak terbaca. oleh karna itu function recover di letakkan pada function endApp dimana function endApp telah di defer yang menyebabkan function endApp akan tetap di jalankan di akhir program walaupun teradi error atau pemberhentian paksa dari function panic, sehingga function recover  tetap terbaca dan bisa membatalkan pemberhentian paksa dari function panic.
	if message != nil {
		fmt.Println("terjadi error", message)
	}

}

func runApp(error bool, value int) {
	defer endApp()
	if error {
		// test := 10 / value
		// fmt.Println(test) // recover juga bisa  menangkap value error dan membatalkan pemberhentian program akibat terjadi error
		panic("APLIKASI ERROR") // panic digunakan untuk menghentikan aplikasi yang secara otomatis script di yang berada dibawah panic tidak akan terekusi karna di paksa berhenti dengan function panic, namun jika ada defer function maka defer function akan tetap dijalankan. fungsi ini seperti throw di bahasa lain
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApplication(10)
	runApp(true, 0)

	fmt.Println("akhir eksekusi")
}
