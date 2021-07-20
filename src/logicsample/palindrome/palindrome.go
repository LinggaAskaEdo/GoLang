package palindrome

import (
	"fmt"
	"strings"
)

func IsPalindrome() {
	var dataLength int
	var data1, data2, allData string

	fmt.Print("Enter first string: ")
	fmt.Scanln(&data1)

	fmt.Print("Enter second string: ")
	fmt.Scanln(&data2)

	allData = data1 + data2
	dataLength = len(allData)

	if len(allData)%2 == 0 {
		split := []rune(allData)
		splitData1 := string(split[0 : dataLength/2])
		splitData2 := string(split[dataLength/2 : dataLength])
		fmt.Printf("%s - %s\n", splitData1, splitData2)
		result := strings.EqualFold(splitData1, reverse(splitData2))

		if result {
			fmt.Println("Palindrome")
		} else {
			fmt.Println("Not Palindrome")
		}
	} else {
		fmt.Println("Not Palindrome")
	}
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
