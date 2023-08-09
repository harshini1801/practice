package main

import "fmt"

func main() {
	greet()
	greet()
	greetby("harshini")
	greetbym("hey", "dileep")
}
func greet() {
	fmt.Println("hello world")
}
func greetby(name string) {
	fmt.Println(name)
}
func greetbym(name, msg string) {
	fmt.Println(name, msg)
}
func areaofrect(l, b float32) float32 { //here float32 is an unamed return type
	return l * b
}
func rect(l, b float32) (area, perimeter float32) {
	area = l * b
	perimeter = 2 * (l + b)
	return area, perimeter
}
