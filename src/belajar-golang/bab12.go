package main

import "fmt"

func main() {
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	var point2 = 8840.0

	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var point3 = 6

	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var point4 = 6

	switch point4 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var point5 = 6

	switch point5 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	var point6 = 6

	switch {
	case point6 == 8:
		fmt.Println("perfect")
	case (point6 < 8) && (point6 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	var point7 = 6

	switch {
	case point7 == 8:
		fmt.Println("perfect")
	case (point7 < 8) && (point7 > 3):
		fmt.Println("awesome")
		fallthrough
	case point7 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	var point8 = 10

	if point8 > 7 {
		switch point8 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point8 == 5 {
			fmt.Println("not bad")
		} else if point8 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point8 == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
