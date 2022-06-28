package main

import (
	"crypto/rand"
)

func randomString(length int) string {
	//rand.Seed(time.Now().UnixNano())
	//b := make([]byte, length)
	//rand.Read(b)
	//
	//return strings.ToUpper(fmt.Sprintf("%x", b)[:length])

	//randomBytes := make([]byte, 32)
	//_, err := rand.Read(randomBytes)
	//if err != nil {
	//	panic(err)
	//}
	//return base32.StdEncoding.EncodeToString(randomBytes)[:length]

	//var dictionary string

	dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var bytes = make([]byte, 6)
	_, err := rand.Read(bytes)
	if err != nil {
		return "012ABC" // set default value, don't worry if duplicate, it will recall to generate again
	}

	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}

	return string(bytes)
}

func main() {
	println(randomString(6))
}
