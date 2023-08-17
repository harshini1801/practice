package main

import (
	"fmt"
	"runtime"
)

func main() {
	go fmt.Println("hello")
	go func() {
		for i := 2; i <= 1000; i++ {
			go evenorodd(i)
		}
	}()
	runtime.Goexit()
}

func evenorodd(num int) {
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
