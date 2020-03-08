package main

import "fmt"

func main() {
	// Inisialisasi Slice
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	fmt.Println("-------------")

	// Hubungan Slice Dengan Array & Operasi Slice
	var fruits2 = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits2[0:2]

	fmt.Println(newFruits)

	fmt.Println("-------------")

	// Slice Merupakan Tipe Data Reference
	var fruits3 = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits3[0:3]
	var bFruits = fruits3[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits3)  // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits3)  // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	fmt.Println("-------------")

	// Fungsi len()
	var fruits4 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits4)) // 4

	fmt.Println("-------------")

	// Fungsi cap()
	var fruits5 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits5)) // len: 4
	fmt.Println(cap(fruits5)) // cap: 4

	var aFruits2 = fruits5[0:3]
	fmt.Println(len(aFruits2)) // len: 3
	fmt.Println(cap(aFruits2)) // cap: 4

	var bFruits2 = fruits5[1:4]
	fmt.Println(len(bFruits2)) // len: 3
	fmt.Println(cap(bFruits2)) // cap: 3

	fmt.Println("-------------")

	// Fungsi append()
	var fruits6 = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits6, "papaya")

	fmt.Println(fruits6) // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	// Sample 2 contoh append
	var fruits7 = []string{"apple", "grape", "banana"}
	var bFruits3 = fruits7[0:2]

	fmt.Println(cap(bFruits3)) // 3
	fmt.Println(len(bFruits3)) // 2

	fmt.Println(fruits7)  // ["apple", "grape", "banana"]
	fmt.Println(bFruits3) // ["apple", "grape"]

	var cFruits2 = append(bFruits3, "papaya")

	fmt.Println(fruits7)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruits3) // ["apple", "grape"]
	fmt.Println(cFruits2) // ["apple", "grape", "papaya"]

	// Fungsi copy()
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	dst2 := []string{"potato", "potato", "potato"}
	src2 := []string{"watermelon", "pinnaple"}
	n2 := copy(dst2, src2)

	fmt.Println(dst2) // watermelon pinnaple potato
	fmt.Println(src2) // watermelon pinnaple
	fmt.Println(n2)   // 2

	// Pengaksesan Elemen Slice Dengan 3 Indeks
	var fruits8 = []string{"apple", "grape", "banana"}
	var aFruits3 = fruits8[0:2]
	var bFruits4 = fruits8[0:2:2]

	fmt.Println(fruits8)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits8)) // len: 3
	fmt.Println(cap(fruits8)) // cap: 3

	fmt.Println(aFruits3)      // ["apple", "grape"]
	fmt.Println(len(aFruits3)) // len: 2
	fmt.Println(cap(aFruits3)) // cap: 3

	fmt.Println(bFruits4)      // ["apple", "grape"]
	fmt.Println(len(bFruits4)) // len: 2
	fmt.Println(cap(bFruits4)) // cap: 2
}
