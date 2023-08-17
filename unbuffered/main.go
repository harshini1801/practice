package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 100
	fmt.Println(<-ch)
}
