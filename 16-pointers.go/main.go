package main

import "fmt"

func main() {
	var num1 int = 100
	var ptr1 *int = &num1
	var ptr2 **int = &ptr1
	var ptr3 *int //called as nill pointer
	fmt.Println(num1)
	fmt.Println(&num1, ptr1, *ptr1, **ptr2)

}
