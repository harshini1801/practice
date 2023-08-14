package main

import "fmt"

var arr1 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

func rowsum(arr1 [3][3]int) [3]int {
	sum := 0
	var arr [3]int
	for i, v := range arr1 {
		for _, v1 := range v {
			sum = sum + v1
		}
		arr[i] = sum
		sum = 0
	}
	return arr
}
func column(arr [3][3]int) [3]int {
	sum := 0
	var arr2 [3]int
	for i, _ := range arr {
		for j, _ := range arr {
			sum = sum + arr[j][i]
		}
		arr2[i] = sum
		sum = 0
	}
	return arr2

}
func main() {
	k := rowsum(arr1)
	fmt.Println(k)
	k1 := column(arr1)
	fmt.Println(k1)
}
