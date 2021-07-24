package palindrometrack

import (
	"fmt"
)

func PalindromeTracking() {
	var data1, data2 string

	fmt.Print("Enter first string: ")
	fmt.Scan(&data1)

	fmt.Print("Enter second string: ")
	fmt.Scan(&data2)

	// data2Length := len(data2)

	// dataArr1 := strings.Split(data1, "")
	// dataArr2 := strings.Split(data2, "")

	// status := false

	for i := 0; i < len(data1); i++ {

	}
	for pos, char := range data2 {
		fmt.Printf("Data %d: %c\n", pos, char)

	}
}
