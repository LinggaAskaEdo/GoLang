package main

import (
	"fmt"
)

type Data struct {
	Name  string
	Grade int
}

func main() {
	// var contactNumber string
	customerContactNumber := "+62085715025257"

	if customerContactNumber[0:4] == "+628" {
		// customerContactNumber = customerContactNumber[1:]
		fmt.Println(customerContactNumber[1:])
	} else {
		fmt.Println("Invalid Number")
	}

	// var name string
	// var grade
	// var breakfast string
	var dataObj Data

	data := map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": nil,
	}

	vName, foundName := data["name"]
	vGrade, foundGrade := data["grade"]
	_, foundBreakfast := data["breakfast"]
	if foundName && foundGrade && foundBreakfast {
		// breakfast = fmt.Sprintf("%v", vBreakfast)
		// fmt.Println(breakfast)

		// fmt.Println(vBreakfast.(string))

		// dataObj = Data{
		// 	Name:  vName.(string),
		// 	Grade: vGrade.(int),
		// 	// Breakfast: sqlx.NullString{String: fmt.Sprintf("%v", vBreakfast), Valid: true},
		// }

		if vName != nil {
			dataObj.Name = vName.(string)
		}

		if vGrade != nil {
			dataObj.Grade = vGrade.(int)
		}
	}

	fmt.Println(dataObj)

	// if reflect.TypeOf(dataObj.Breakfast) == nil {
	// 	fmt.Println("ZZZ")
	// }

	// if dataObj.Breakfast.String == "" {
	// 	fmt.Println("AAA")
	// } else {
	// 	fmt.Println("BBB")
	// }

	// fmt.Println("Hello, Reader! Your learning about 'goto' statement")
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// // We create a for loop which runs until i is 10
	// for _, fruit := range fruits {
	// 	for i := 0; i < 4; i++ {
	// 		proc1(fruit)
	// 		proc2(fruit)
	// 	}
	// }
}
