package main

func main() {
	a := 100
	b := 100
	k1 := add(a, b)
	println(k1)
}
func add(a1, b1 any) int {
	return a1.(int) + b1.(int)
}
