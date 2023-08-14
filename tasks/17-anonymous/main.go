package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello secret")
	}()
	func(str string) {
		fmt.Println(str)
	}("hello welcome to victoria")
	k1, k2, k3 := func(str string) (int, int, int) {
		count1 := 0
		count2 := 0
		count3 := 0
		for _, v := range str {
			switch v {
			case 97, 105, 101, 111, 117:
				{
					count1 = count1 + 1
				}
			case 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 58, 59, 60, 61, 62, 63, 64, 123, 124, 125, 126, 91, 92, 93, 94, 95, 96:
				{
					count2 = count2 + 1
				}
			default:
				count3 = count3 + 1
			}
		}
		return count1, count2, count3
	}("harshi")
	fmt.Println(k1, k2, k3)
}
