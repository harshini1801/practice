package main

import "fmt"

func main() {
	i := 1
loop:
	if i%2 == 0 {
		fmt.Println(i, "is even")

	}
	i++
	if i <= 10 {
		goto loop
	}

}
