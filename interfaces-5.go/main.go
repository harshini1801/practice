package main

import "fmt"

type icalculator interface {
	add(int) icalculator
	get() int
	sub(int) icalculator
}
type inttype struct {
	result int
}

func (i *inttype) add(num int) icalculator {
	i.result = i.result + num
	return i
}
func (i *inttype) sub(num int) icalculator {
	i.result = i.result - num
	return i
}
func (i *inttype) mul(num int) icalculator {
	if i.result == 0 {
		i.result = 1
	}
	i.result = i.result * num
	return i
}

func (i *inttype) get() int {
	return i.result
}
func main() {
	var icalc icalculator

	icalc = new(inttype)

	r1 := icalc.add(10).add(20).add(20).add(30).add(30).get()
	fmt.Println("Result r1:", r1)

	icalc = new(inttype)
	r2 := icalc.add(10).sub(20).add(20).get()
	fmt.Println("Result r2:", r2)

}
