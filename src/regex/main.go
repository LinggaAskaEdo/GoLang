package main

import (
	"log"
	"regexp"
)

var (
	regexList = []string{"^(.*[a-z]+.*)$", "^(.*[A-Z]+.*)$", "^(.*[!@#$&]+.*)$", "^(.*[0-9]+.*)$", "^.{8,}$"}
)

func main() {
	dataList := []string{"T1t!", "Tt1", "Tit1T", "tT", "tt", "tt1", "TT", "TT1", "tT1", "t1T1", "Pas$w0rd", "P@ssword123", "Ar$tdhnei0", "~1A#$%^n*()_+", "new_P@ssw0rd"}
	//regexList := []string{"^(.*[a-z]+.*)$", "^(.*[A-Z]+.*)$", "^(.*[!@#$&]+.*)$", "^(.*[0-9]+.*)$", "^.{8,}$"}
	//regexList := []string{"^([a-z]+.*)$", "^([A-Z]+)$", "^([!@#$&]+)$", "^([0-9]+)$", "^.{8,}$"}

	//regex := "^(\\A[a-z]+\\z)(\\A[A-Z]+\\z)(\\A[0-9]+\\z)(\\A[!@#$&]+\\z)(\\A.{8,}\\z)$"
	//regex := "(^.*[0-9]+.*$)(^.*[a-z]+.*$)"

	//regex := "^(?=.{8,}$)(?=.[A-Z])(?=.[a-z])(?=.[0-9])(?=.\\W).*$"

	//regPass := "(.*[A-Za-z]+.*)(.*[!@#$&]+.*)(.*[0-9]+.*)"
	//regPass1 := "^(.*[A-Za-z]+.*)$"
	//regPass2 := "^(.*[!@#$&]+.*)$"
	//regPass3 := "^(.*[0-9]+.*)$"
	//regPass4 := "^.{3,}$"
	//regPass := "^([^\\s][A-Za-z])([^\\s][!@#$&])([^\\s]*[0-9]).{8,}$"
	//regPass := "^(?=.{8,}$)(?=.[A-Z])(?=.[a-z])(?=.[0-9])(?=.\\W).*$"
	//regPass := "^.?[a-z].?[A-Z].?[0-9].?[!@#$&].?$"
	//regPass := "^(.*?[a-z].*?)(.*?[A-Z].*?)(.*?[0-9].*?)$"
	//regPass := "^(.*)([a-z]{1,}?)(.*)([0-9]{1,}?)(.*)([A-Z]{1,}?)(.*)$"
	//reg1 := "^[a-z]{1,}?$"
	//reg2 := "^[A-Z]{1,}?$"
	//reg3 := "^[0-9]{1,}?$"
	//regPass := "[a-z]{1,}?"

	//re := regexp.MustCompile(regex)
	//re :=  regexp.MustCompile(`^[a-z]+$",,"NEW NAME\(S\)(.*)"`)
	//fmt.Println("Matched:", re.MatchString("pass"))

	status := true
	for _, data := range dataList {
		for i, regex := range regexList {
			status = true
			result, err := regexp.MatchString(regex, data)
			if err != nil {
				log.Println("Error: ", err)
			}

			if !result {
				log.Println("invalid regex ", i+1)
				status = false
				break
			}
		}

		if status {
			log.Println(data, ": valid")
		} else {
			log.Println(data, ": invalid")
		}
	}

	//for _, data := range dataList {
	//	status, err := regexp.MatchString(regex, data)
	//	if err != nil {
	//		log.Println("Error: ", err)
	//	}
	//
	//	if status {
	//		log.Println(data, ": valid")
	//	} else {
	//		log.Println(data, ": invalid")
	//	}
	//}

	//sampleRegex := regexp.MustCompile(regPass1 + regPass2 + regPass3 + regPass4)

	//for _, data := range dataList {
	//log.Println(data, ": ", sampleRegex.Match([]byte(data)))

	//result1, err := regexp.MatchString(regPass1, data)
	//if err != nil {
	//	log.Println("Error: ", err)
	//	return
	//}
	//
	//result2, err := regexp.MatchString(regPass2, data)
	//if err != nil {
	//	log.Println("Error: ", err)
	//	return
	//}
	//
	//result3, err := regexp.MatchString(regPass3, data)
	//if err != nil {
	//	log.Println("Error: ", err)
	//	return
	//}
	//
	//result4, err := regexp.MatchString(regPass4, data)
	//if err != nil {
	//	log.Println("Error: ", err)
	//	return
	//}
	//
	//if result1 {
	//	log.Println(data, ": valid1")
	//} else {
	//	log.Println(data, ": invalid1")
	//}
	//
	//if result2 {
	//	log.Println(data, ": valid2")
	//} else {
	//	log.Println(data, ": invalid2")
	//}
	//
	//if result3 {
	//	log.Println(data, ": valid3")
	//} else {
	//	log.Println(data, ": invalid3")
	//}
	//
	//if result4 {
	//	log.Println(data, ": valid4")
	//} else {
	//	log.Println(data, ": invalid4")
	//}
	//
	//if result1 && result2 && result3 && result4 {
	//	log.Println(data, ": valid")
	//} else {
	//	log.Println(data, ": invalid")
	//}
	//}
}
