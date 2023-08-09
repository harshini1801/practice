package main

import (
	"fmt"
	"reflect"
)

func main() {
	var age uint8 = 38
	fmt.Println(age)
	fmt.Println(reflect.TypeOf(age))
	var num1 = 100 //another kind of declaration
	num2 := 200    //short hand declaration
	fmt.Println(num1, num2)
	//type inference
	var (
		num3 int
		num4 float32
		ok1  bool
		str1 string
	)
	fmt.Println(num3, num4, ok1, str1)

}
