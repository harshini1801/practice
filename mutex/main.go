package main

import (
	"fmt"
	"runtime"
)

var num int16

func main() {
	notify := make(chan struct{})
	go func() {
		for i := 1; i <= 1000; i++ {
			go incr()
			notify <- struct{}{}
		}
	}()
	fmt.Println(num)
	runtime.Goexit()

}
func incr() {
	num++
	fmt.Println(num)
}
