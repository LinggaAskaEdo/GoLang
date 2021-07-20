package main

import (
	"fmt"
	"logicsample/anagram"
	"logicsample/largestelement"
	"logicsample/palindrome"
	"logicsample/palindromenumber"
	"logicsample/palindrometrack"
	"logicsample/separatewrapper"
	"logicsample/wrapper"
)

func main() {
	var moduleNumber int

	fmt.Print("Enter the module number: ")
	fmt.Scanln(&moduleNumber)

	switch moduleNumber {
	case 1:
		largestelement.FindLargestElementInArray()
	case 2:
		palindrome.IsPalindrome()
	case 3:
		palindromenumber.IsPalindromeNumber()
	case 4:
		anagram.IsAnagram()
	case 5:
		palindrometrack.PalindromeTracking()
	case 6:
		wrapper.WrapperTest()
	case 7:
		separatewrapper.SeparateWrapperTest()
	default:
		fmt.Println("Undefine module number")
	}
}
