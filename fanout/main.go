package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := multiple(ch1, ch2)
	go func() {
		ch1 <- "victroia"
		ch2 <- "secrets &co"

	}()
	for val := range ch3 {
		fmt.Println(val)
	}

}
func multiple(ch1, ch2 chan string) chan string {
	ch := make(chan string)
	go func() {
		ch <- <-ch1 + <-ch2
		close(ch)

	}()
	return ch

}
