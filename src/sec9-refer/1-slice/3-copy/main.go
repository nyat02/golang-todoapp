package main

import "fmt"

func main() {
	sl := []int{100, 200}
	sl2 := sl
	sl2[0] = 1000

	fmt.Println(sl, sl2)

	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 10)
	fmt.Println(sl3, sl4)
	n := copy(sl4, sl3)
	sl4[0] = 100
	fmt.Println(sl3, sl4, n)
}
