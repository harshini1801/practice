package main

import "fmt"

func sort(arr1 []int) []int {
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[i] > arr1[j] {
				arr1[i], arr1[j] = arr1[j], arr1[i]
			}
		}
	}
	return arr1

}
func main() {
	fmt.Printf("Hello, World!")
	arr := []int{5, 7, 3, 8, 2, 2, 4, 6, 7, 7}
	arr = sort(arr)
	fmt.Println(arr)
	secondbig := arr[len(arr)-2]
	secondsmall := arr[1]
	fmt.Println(secondbig, secondsmall)

}
