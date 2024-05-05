package main

import "fmt"

func Plus(x int, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, _ := Div(9, 4) // 返り値を片方破棄し、未使用エラーを回避
	fmt.Println(i2)

	i4 := Double(100)
	fmt.Println(i4)

	noreturn()
}
