package main

import (
	"log"
)

func main() {
	//
	// dateNow := time.Now().Format(time.RFC3339)
	// dateNowUTC := time.Now().UTC().Format(time.RFC3339)

	// fmt.Println("dateNow: ", dateNow, ", dateNowUTC: ", dateNowUTC)

	// phone := "85715025257"
	// if phone[0] != '8' {
	// 	log.Println("AAA")
	// } else {
	// 	log.Println("BBB")
	// }

	//
	// result := ""
	// input1 := "edo"

	// for i := len(input1) - 1; i >= 0; i-- {
	// 	result += string(input1[i])
	// }

	// log.Println(result)

	//
	input2 := []int{8, 2, 7, 13}

	log.Println(drop(input2, 2))
}

func drop(input []int, counter int) []int {
	log.Println("length: ", len(input), ", counter: ", counter)
	// if counter == 0 {
	// 	return input
	// } else if len(input) < counter {
	// 	return nil
	// }
	log.Println("AAA")
	result := make([]int, 0, len(input)-counter)

	log.Println("BBB")
	log.Println(input[counter])
	result = append(result, input[counter])
	log.Println(result)
	log.Println("CCC")
	counter++

	if counter != len(input) {
		log.Println("DDD")
		drop(input, counter)
	}

	return result
}
