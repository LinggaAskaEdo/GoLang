package main

import "log"

func main() {
	status := true
	num := 0

	for ok := true; ok; ok = status {
		log.Println(num)
		num++
		if num == 3 {
			status = false
		} else if num == 2 {
			return
		}
	}
}
