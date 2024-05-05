package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2)

	fl := 10.5
	i3 := int(fl)
	fmt.Printf("i3 = %T\n", i3)
	fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i4, _ := strconv.Atoi(s)
	fmt.Println(i4)
	fmt.Printf("i4 = %T\n", i4)

	i5, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i5)
	fmt.Printf("i5 = %T\n", i5)

	var i6 int = 200
	s2 := strconv.Itoa(i6)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}
