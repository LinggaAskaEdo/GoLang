package largestelement

import "fmt"

func FindLargestElementInArray() {
	var n [100]float32
	var total int

	fmt.Print("Enter number of elements: ")

	fmt.Scanln(&total)

	for i := 0; i < total; i++ {
		fmt.Print("Enter the number: ")
		fmt.Scan(&n[i])
	}

	for j := 1; j < total; j++ {
		if n[0] < n[j] {
			n[0] = n[j]
		}
	}

	fmt.Println("The largest number is: ", n[0])
}
