package palindromenumber

import "fmt"

func IsPalindromeNumber() {
	var number, remainder, temp int
	var reverse int = 0

	fmt.Print("Enter any positive integer: ")
	fmt.Scan(&number)

	temp = number

	fmt.Printf("remainder: %d\n", remainder)
	fmt.Printf("reverse: %d\n", reverse)
	fmt.Printf("number: %d\n", number)

	// For Loop used in format of While Loop
	for {
		remainder = number % 10
		fmt.Printf("remainder: %d\n", remainder)
		reverse = reverse*10 + remainder
		fmt.Printf("reverse: %d\n", reverse)
		number /= 10
		fmt.Printf("number: %d\n", number)

		if number == 0 {
			break // Break Statement used to exit from loop
		}
	}

	if temp == reverse {
		fmt.Printf("%d is a Palindrome\n", temp)
	} else {
		fmt.Printf("%d is not a Palindrome\n", temp)
	}
}
