package main

import "fmt"

func main() {
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	n2 := 6
	switch {
	case n2 > 0 && n2 < 4:
		fmt.Println("0 < n2 < 4")
	case n2 > 3 && n2 < 6:
		fmt.Println("3 < n2 < 6")
	default:
		fmt.Println("I don't know")
	}
}
