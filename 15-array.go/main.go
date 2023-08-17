package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{6, 7, 8, 9, 10}
	//var arr3 [5]int = arr1
	//var arr4 [4]int=arr2//this gives errorr//
	//multidimensional array
	fmt.Println(arr1)
	fmt.Println(arr2)
	arr2d1 := [2][2]int{{1, 2}, {3, 4}}
	for _, arr1d := range arr2d1 {
		for _, v := range arr1d {
			fmt.Print(v, " ")
		}
	}
	for i=0
}
