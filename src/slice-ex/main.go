package main

import "log"

func main() {
	whitelist := map[int64]bool{
		13: true,
		66: false,
	}

	testList := []int64{22, 66}

	for _, test := range testList {
		log.Print(test)
		if whitelist[test] {
			log.Println("Ada")
		} else {
			log.Println("Ga Ada")
		}
	}
}
