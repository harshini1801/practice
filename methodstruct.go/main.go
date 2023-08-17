package main

import "fmt"

type rect struct {
	l, b float32
}
type square struct {
	s float32
}

func areaofrect(r rect) float32 {
	return r.l * r.b

}
func areaofsquare(s1 square) float32 {
	return s1.s * s1.s
}

// to write method need to add a receiver
func (r rect) Area() float32 {
	return r.l * r.b
}
func main() {
	r1 := rect{l: 1.2, b: 3.2}
	k := r1.Area()
	fmt.Println(k)
}
