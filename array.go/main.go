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
}
