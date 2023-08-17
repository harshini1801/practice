package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println(arr1)
	var arr2 [5]bool
	fmt.Println(arr2)
	var arr3 [3]any
	fmt.Println(arr3)
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5
	fmt.Println(len(arr1))
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])

	}
	//for range loop
	//range
	k := sumofarr(arr1)
	fmt.Println(k)

}
func sumofarr(arr [5]int) int {
	sum := 0
	for _, value := range arr {
		sum = sum + value
	}
	return sum
}
