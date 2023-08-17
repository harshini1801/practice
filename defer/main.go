package main

import "fmt"

func main() {
	func() {
		defer fmt.Println("hii")
		fmt.Println("hey")
	}()
	defer fmt.Println("this is end of the main")
	defer fmt.Println("hello victoria")
	defer fmt.Println("hello secret")
	fmt.Println("welcome to the brand")

}
