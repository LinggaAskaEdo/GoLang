package main

import "fmt"

func main() {
	// Penggunaan Map
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// Inisialisasi Nilai Map
	// cara vertikal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara horizontal
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	fmt.Println(chicken1)
	fmt.Println(chicken2)

	// Iterasi Item Map Menggunakan for - range
	var chicken3 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    70,
	}

	for key, val := range chicken3 {
		fmt.Println(key, " \t:", val)
	}

	// Menghapus Item Map
	var chicken4 = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken4)) // 2
	fmt.Println(chicken4)

	delete(chicken4, "januari")

	fmt.Println(len(chicken4)) // 1
	fmt.Println(chicken4)

	// Deteksi Keberadaan Item Dengan Key Tertentu
	var chicken5 = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken5["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// Kombinasi Slice & Map
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
