package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Brad"
	var age int32 = 31
	const isCool = false

	// Shorthand
	// name := "Aska"
	// email := "a5k4_fansarsenal@yahoo.com"
	size := 1.3

	name, email, office := "Aska", "a5k4_fansarsenal@yahoo.com", "Plaza Oleos"

	fmt.Println(name, "-", age, "-", isCool, "-", email, "-", office)
	fmt.Printf("%T\n", size)
}
