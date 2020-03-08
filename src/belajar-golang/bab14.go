package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	// Inisialisasi Nilai Array
	fmt.Println("---------------------")

	// cara horizontal
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	fmt.Println("---------------------------------------------------")

	var fruits2 [4]string

	// cara vertikal
	fruits2 = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t\t", len(fruits2))
	fmt.Println("Isi semua element \t", fruits2)

	// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	fmt.Println("-----------------------------")

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("jumlah elemen \t:", len(numbers))
	fmt.Println("data array \t:", numbers)

	// Array Multidimensi
	fmt.Println("-----------------------------")

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)
}
