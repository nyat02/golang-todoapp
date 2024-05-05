package main

import "fmt"

const Pi = 3.14

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 2
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

// var Big int = 9223372036854775807 + 1
// const Big int = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)

	// Pi = 3 //不可
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	// fmt.Println(Big - 1)
	// 講義通りならば、以下のエラーが出ないはずだが、出る

	fmt.Println(c0, c1, c2)
}
