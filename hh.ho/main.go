package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	r := s
	r = []int{1, 4, 5, 6}
	fmt.Println(s, r)
}
