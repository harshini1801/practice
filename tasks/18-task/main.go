package main

import (
	"fmt"
)

func main() {

	defer recoverme()

	sum1, sub1, mul1, div1 := Calc("Hi", "Vitoria") // error panic

	fmt.Println("sum:", sum1, "sub:", sub1, "mul:", mul1, "div:", div1)

	sum2, sub2, mul2, div2 := Calc([]int{10, 12}, 10) // error panic

	fmt.Println("sum:", sum2, "sub:", sub2, "mul:", mul2, "div:", div2)

	sum3, sub3, mul3, div3 := Calc(10, 20) // should give 30,-10,200,0

	fmt.Println("sum:", sum3, "sub:", sub3, "mul:", mul3, "div:", div3)

}

func Calc(a, b any) (float64, float64, float64, float64) {

	var sum, sub, mul, div float64

	defer recoverme()

	switch a.(type) { // this is how you need to check type of a variable which is of type any

	case int:

		{

			sum = float64(a.(int) + b.(int))

			sub = float64(a.(int) - b.(int))

			mul = float64(a.(int) * b.(int))

			div = float64(a.(int) / b.(int))

		}

	case uint:

		{

			sum = float64(a.(uint) + b.(uint))

			sub = float64(a.(uint) - b.(uint))

			mul = float64(a.(uint) * b.(uint))

			div = float64(a.(uint) / b.(uint))

		}

	case uint8:

		{

			sum = float64(a.(uint8) + b.(uint8))

			sub = float64(a.(uint8) - b.(uint8))

			mul = float64(a.(uint8) * b.(uint8))

			div = float64(a.(uint8) / b.(uint8))

		}

	case uint16:

		{

			sum = float64(a.(uint16) + b.(uint16))

			sub = float64(a.(uint16) - b.(uint16))

			mul = float64(a.(uint16) * b.(uint16))

			div = float64(a.(uint16) / b.(uint16))

		}

	case uint32:

		{

			sum = float64(a.(uint32) + b.(uint32))

			sub = float64(a.(uint32) - b.(uint32))

			mul = float64(a.(uint32) * b.(uint32))

			div = float64(a.(uint32) / b.(uint32))

		}

	case uint64:

		{

			sum = float64(a.(uint64) + b.(uint64))

			sub = float64(a.(uint64) - b.(uint64))

			mul = float64(a.(uint64) * b.(uint64))

			div = float64(a.(uint64) / b.(uint64))

		}

	case int8:

		{

			sum = float64(a.(int8) + b.(int8))

			sub = float64(a.(int8) - b.(int8))

			mul = float64(a.(int8) * b.(int8))

			div = float64(a.(int8) / b.(int8))

		}

	case int16:

		{

			sum = float64(a.(int16) + b.(int16))

			sub = float64(a.(int16) - b.(int16))

			mul = float64(a.(int16) * b.(int16))

			div = float64(a.(int16) / b.(int16))

		}

	case int32:

		{

			sum = float64(a.(int32) + b.(int32))

			sub = float64(a.(int32) - b.(int32))

			mul = float64(a.(int32) * b.(int32))

			div = float64(a.(int32) / b.(int32))

		}

	case int64:

		{

			sum = float64(a.(int64) + b.(int64))

			sub = float64(a.(int64) - b.(int64))

			mul = float64(a.(int64) * b.(int64))

			div = float64(a.(int64) / b.(int64))

		}

	case float32:

		{

			sum = float64(a.(float32) + b.(float32))

			sub = float64(a.(float32) - b.(float32))

			mul = float64(a.(float32) * b.(float32))

			div = float64(a.(float32) / b.(float32))

		}

	case float64:

		{

			sum = a.(float64) + b.(float64)

			sub = a.(float64) - b.(float64)

			mul = a.(float64) * b.(float64)

			div = a.(float64) / b.(float64)

		}

	default:

		panic("Invalid data types")

	}

	return sum, sub, mul, div

}

func recoverme() {

	if p := recover(); p != nil {

		fmt.Println("Panic recovered")

	}

}
