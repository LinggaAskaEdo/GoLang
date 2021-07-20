package anagram

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func IsAnagram() {
	var data1, data2 string

	fmt.Print("Enter first string: ")
	fmt.Scanln(&data1)

	fmt.Print("Enter second string: ")
	fmt.Scanln(&data2)

	if len(data1) != len(data2) {
		fmt.Printf("%s and %s not anagrams string\n", data1, data2)
	} else {
		// arr1 := []rune(data1)
		// arr2 := []rune(data2)
		// arr1 := strings.Fields(data1)
		// arr2 := strings.Fields(data2)
		arr1 := strings.Split(data1, "")
		arr2 := strings.Split(data2, "")

		fmt.Printf("arr1: %s\n", arr1)
		fmt.Printf("arr2: %s\n", arr2)

		// sort.Sort(sort.StringSlice(arr1))
		// sort.Sort(sort.StringSlice(arr2))

		sort.Strings(arr1)
		sort.Strings(arr2)

		fmt.Printf("arr1: %s\n", arr1)
		fmt.Printf("arr2: %s\n", arr2)

		if reflect.DeepEqual(arr1, arr2) {
			fmt.Printf("%s and %s anagrams string\n", data1, data2)
		} else {
			fmt.Printf("%s and %s not anagrams string\n", data1, data2)
		}
	}
}
