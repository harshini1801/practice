package main

import "fmt"

func main() {
	arr := []int{5, 8, 9, 4, 3, 2, 121, 45, 142, 190, 7, 6, 4, 3, 2, 200, 3, 8}
	k1, k2 := bigsmall(arr)
	fmt.Println(k1, k2)

}
func bigsmall(arr []int) (int, int) {
	big := arr[0]
	small := arr[0]
	for i := 1; i < len(arr); i++ {
		if big < arr[i] {
			big = arr[i]
		}

	}
	for j := 1; j < len(arr); j++ {
		if small > arr[j] {
			small = arr[j]
		}
	}
	return big, small

}
