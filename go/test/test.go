package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Scan(&x, &y)

	s := x * y
	c := (x + y) * 2

	fmt.Println(s, c)
}
