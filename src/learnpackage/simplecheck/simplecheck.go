package simplecheck

import "fmt"

// If else statement
func SimpleCheck1() {
	num := 11

	if num%2 == 0 { //checks if number is even
		fmt.Println("The number", num, "is even")
	} else {
		fmt.Println("The number", num, "is odd")
	}
}

// If ... else if ... else statement
func SimpleCheck2() {
	num := 99

	if num <= 50 {
		fmt.Println(num, "is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println(num, "is between 51 and 100")
	} else {
		fmt.Println(num, "is greater than 100")
	}
}

// If with assignment
func SimpleCheck3() {
	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
