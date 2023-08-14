package main

import "fmt"

func main() {

	var num1 int = 64
	var num2 int8 = 32
	var num3 int16 = 81
	var num4 int32 = 88
	var num5 int64 = 225
	var num6 uint = 233
	var num7 uint8 = 225
	var num8 uint16 = 276
	var num9 uint32 = 255
	var num10 uint64 = 87
	var num11 float32 = 64
	var num12 float64 = 78
	var num13 byte = 'C'
	var num14 rune = 'B'
	var iface1 any
	iface1 = 100
	var iface2 any
	iface2 = 32.2
	var result1 float64
	ar result2 float64
	result1 = float64(num1) + float64(num2) + float64(num3) + float64(num4) + float64(num5) + float64(num6) + float64(num7) + float64(num8) + float64(num9) + float64(num10) + float64(num11) + num12 + float64(num13) + float64(num14) + float64(iface1.(int)) + iface2.(float64)
	fmt.Println(result1)
	

}
