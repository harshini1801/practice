package main

import "fmt"

type icalculator interface {
	add(any) icalculator
	get() float64
	sub(any) icalculator
}
type intype struct {
	result float64
}

func (i *intype) add(num any) icalculator {
	switch num.(type) { // this is how you need to check type of a variable which is of type any
	case int:
		i.result = i.result + float64(num.(int))
	case uint:
		i.result = i.result + float64(num.(uint))
	case uint8:
		i.result = i.result + float64(num.(uint8))
	case uint16:
		i.result = i.result + float64(num.(uint16))
	case uint32:
		i.result = i.result + float64(num.(uint32))
	case uint64:
		i.result = i.result + float64(num.(uint64))
	case int8:
		i.result = i.result + float64(num.(int8))

	case int16:
		i.result = i.result + float64(num.(int64))
	case int32:
		i.result = i.result + float64(num.(int32))
	case int64:
		i.result = i.result + float64(num.(int64))
	case float32:
		i.result = i.result + float64(num.(float32))
	case float64:
		i.result = i.result + num.(float64)
	}
	return i
}
func (i *intype) sub(num any) icalculator {
	switch num.(type) { // this is how you need to check type of a variable which is of type any
	case int:
		i.result = i.result - float64(num.(int))
	case uint:
		i.result = i.result - float64(num.(uint))
	case uint8:
		i.result = i.result - float64(num.(uint8))
	case uint16:
		i.result = i.result - float64(num.(uint16))
	case uint32:
		i.result = i.result - float64(num.(uint32))
	case uint64:
		i.result = i.result - float64(num.(uint64))
	case int8:
		i.result = i.result - float64(num.(int8))

	case int16:
		i.result = i.result - float64(num.(int64))
	case int32:
		i.result = i.result - float64(num.(int32))
	case int64:
		i.result = i.result - float64(num.(int64))
	case float32:
		i.result = i.result - float64(num.(float32))
	case float64:
		i.result = i.result - num.(float64)
	}
	return i
}

func (i *intype) get() float64 {
	return i.result
}
func main() {
	var icalc icalculator

	icalc = new(intype)

	r1 := icalc.add(10).add(20).add(20).add(30).add(30).get()
	fmt.Println("Result r1:", r1)

	icalc = new(intype)
	r2 := icalc.add(10).sub(20).add(20).get()
	fmt.Println("Result r2:", r2)

}
