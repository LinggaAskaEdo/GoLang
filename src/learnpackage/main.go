package main

import (
	"fmt"
	"learnpackage/simplearraysslices"
	"learnpackage/simplecheck"
	"learnpackage/simpleinterest"
	"learnpackage/simpleloop"
	"learnpackage/simpleswitch"
	"log"
)

var p, r, t = 5000.0, 10.0, 1.0

func init() {
	println("Main package initialized")

	if p < 0 {
		log.Fatal("Principal is less than zero")
	}

	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}

	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

func main() {
	fmt.Println("Simple interest calculation")
	p := 5000.0
	r := 10.0
	t := 1.0
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)

	simplecheck.SimpleCheck1()
	simplecheck.SimpleCheck2()
	simplecheck.SimpleCheck3()

	simpleloop.SimpleLoop1()
	simpleloop.SimpleLoop2()
	simpleloop.SimpleLoop3()
	simpleloop.SimpleLoop4()
	simpleloop.SimpleLoop5()

	simpleswitch.SimpleSwitch1()
	simpleswitch.SimpleSwitch2()
	simpleswitch.SimpleSwitch3()
	simpleswitch.SimpleSwitch4()
	simpleswitch.SimpleSwitch5()
	simpleswitch.SimpleSwitch6()
	simpleswitch.SimpleSwitch7()

	simplearraysslices.SimpleArraysSlices1()
	simplearraysslices.SimpleArraysSlices2()
	simplearraysslices.SimpleArraysSlices3()
	simplearraysslices.SimpleArraysSlices4()
	simplearraysslices.SimpleArraysSlices5()
	simplearraysslices.SimpleArraysSlices6()
	simplearraysslices.SimpleArraysSlices7()
	simplearraysslices.SimpleArraysSlices8()
}
