package main

import "fmt"

func main() {
	a := 10
	b := 20
	defer fmt.Println(a, b, add(a, b))
	defer func(int, int) {
		c := add(a, b)
		fmt.Println(a, b, c)
	}(a, b)
	a = 11
	b = 21
	c := add(a, b)
	fmt.Println(c)

}
func add(a, b int) int {
	return a + b
}
