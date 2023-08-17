package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	r := s
	r = append(s)
	fmt.Println(r)
	fmt.Println(s)
}
