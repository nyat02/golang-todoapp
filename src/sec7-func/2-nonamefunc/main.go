package main

import "fmt"

func main() {
	f := func(x, y int) int {
		return x + y
	}
	i := f(3, 4)
	fmt.Println(i)

	i2 := func(x, y int) int {
		return x + y
	}(3, 4)
	fmt.Println(i2)
}
