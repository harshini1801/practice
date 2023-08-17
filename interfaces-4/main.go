package main

import "fmt"

func main() {

	var icalc icalculator
	if icalc == nil {
		fmt.Println("nil")
	}
	it1 := &inttype{a: 100, b: 200}
	icalc = it1
	k1 := icalc.add()
	k2 := icalc.sub()
	fmt.Println(k1, k2)
	it2 := &float32type{a: 40, b: 50}
	icalc = it2
	k3 := icalc.add()
	k4 := icalc.sub()
	fmt.Println(k3, k4)
	/*
		it3 := &stringtype{str1: "harshi", str2: "dileep"}
		icalc = it3
		k5:= icalc.add()
		fmt.Println(k5)
	*/
	it4 := &float64type{a: 10, b: 20}
	icalc = it4
	k6 := icalc.add()
	k7 := icalc.sub()
	fmt.Println(k6, k7)

}

type icalculator interface {
	add() any
	sub() any
}
type inttype struct {
	a, b   int
	result int
}

func (i *inttype) add() any {
	i.result = i.a + i.b
	return i.result
}
func (i *inttype) sub() any {
	i.result = i.a - i.b
	return i.result
}

type float32type struct {
	a, b   float32
	result float32
}

func (f *float32type) add() any {
	f.result = f.a + f.b
	return f.result
}
func (f *float32type) sub() any {
	f.result = f.a - f.b
	return f.result
}

type stringtype struct {
	str1, str2 string
	result     string
}

func (s *stringtype) add() any {
	s.result = s.str1 + s.str2
	return s.result
}

type float64type struct {
	a, b   float64
	result float64
}

func (f4 *float64type) sub() any {
	f4.result = f4.a - f4.b
	return f4.result

}
func (f4 *float64type) add() any {
	f4.result = f4.a + f4.b
	return f4.result

}
func (f *float32type) display() any {
	fmt
}
