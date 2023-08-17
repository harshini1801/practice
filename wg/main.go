package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	// Unbuffered channel has been instantiated

	go func() {
		num := <-ch // receive a value from the channel
		fmt.Println(num)

	}()
	//go func() {
	time.Sleep(time.Second * 5)
	ch <- 100
	// Assign a value to the channel

}
