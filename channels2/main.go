package main

import "fmt"

func main() {
	ch := make(chan int)
	notify := make(chan struct{})
	go func() {
		for i := 0; i <= 100; i++ {
			ch <- i * i
		}
	}()
	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Println("receiver-->", <-ch)
		}
		notify <- struct{}{}
	}()
	<-notify
}
