package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	dateNow := time.Now().Format(time.RFC3339)
	dateNowUTC := time.Now().UTC().Format(time.RFC3339)

	fmt.Println("dateNow: ", dateNow, ", dateNowUTC: ", dateNowUTC)

	phone := "85715025257"
	if phone[0] != '8' {
		log.Println("AAA")
	} else {
		log.Println("BBB")
	}
}
