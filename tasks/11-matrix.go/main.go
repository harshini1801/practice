package main

import "fmt"

func main() {
	var arr1 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := rowtocolum(arr1)
	fmt.Println(k)

}
func rowtocolum(arr [3][3]int) [3][3]int {
	var arr2 [3][3]int

	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2); j++ {
			arr2[i][j] = arr[j][i]
		}
	}
	return arr2
}
