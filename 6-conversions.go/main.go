package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num1 uint8 = 38
	var num2 int = int(num1)
	var num3 int = 65001
	var num4 uint8 = uint8(num3)
	fmt.Println(num2, num4)
	var str1 string = "h"
	var num5 int = int(str1[0])
	println(num5)
	var num6 int = 65
	var str2 string = string(num6)
	fmt.Println(reflect.TypeOf(str2), str2)

}
