package main

import (
	"fmt"
	"reflect"
)

// arrays are stored in stack memory
// we dontb use arrays as parameter
func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Print(reflect.TypeOf(arr))
	str := `"Hello world"`
	fmt.Println(str)
	//swapping in go
	a := 20
	b := 10
	c := 30
	a, b, c = c, a, b //swapping in go

}
