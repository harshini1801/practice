package main

import "fmt"

type integer = int

func main() {
	var num1 integer = 100
	var num2 rune = 'A'
	var num3 int32 = 'B'
	var num4 uint8 = 'A'
	var num5 uint8 = 'B'
	fmt.Println(num1, num2, num3, num4, num5)
}
