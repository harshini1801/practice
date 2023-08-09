package main

import (
	"fmt"
	"reflect"
)

func main() {
	var iface1 any



func add(a, b any) (float64, error) {
	sum:=0.//if you give 0. or 0.0 sum is treated as  float64
	
	if reflect.TypeOf(a) !=reflect.TypeOf(b){
		return 0,errors.New("a and b must be in same type")
	}
	return sum ,nil

}
//use generics-generics are introduced only in 1.18 version
//use switch type
