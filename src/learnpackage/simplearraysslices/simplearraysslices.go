package simplearraysslices

import "fmt"

func SimpleArraysSlices1() {
	var a [3]int //int array with length 3
	a[0] = 12    // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)
}

// short hand declaration
func SimpleArraysSlices2() {
	a := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(a)
}

//ignore length of array
func SimpleArraysSlices3() {
	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a)
}

// Arrays are value types
func SimpleArraysSlices4() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

func SimpleArraysSlices5() {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
}

// Length of an array
func SimpleArraysSlices6() {
	a := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(a))

}

// Iterating arrays using range
func SimpleArraysSlices7() {
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}
}
func SimpleArraysSlices8() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}

// Multidimensional arrays
