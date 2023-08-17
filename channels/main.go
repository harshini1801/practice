package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch1 <- 100


	
	num := <-ch1
	fmt.Println(num)
}
