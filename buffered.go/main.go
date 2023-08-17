package main

import "fmt"

func main() {
	ch := make(chan int, 2) //buffered channel
	ch <- 100
	ch <- 200
	//ch <- 300 //it is blocked because 2 is length

	fmt.Println(<-ch, <-ch)
	ch <- 300
	fmt.Println(<-ch)

}
