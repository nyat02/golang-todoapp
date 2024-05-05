package main

import "fmt"

func integers(start int) func() int {
	i := start - 1
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers(0)
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := integers(100)
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
}
