package main

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	chs := []chan int{ch1, ch2, ch3}
	ch4 := variadichan(chs...)
	go func() {
		ch1 <- 10
		ch2 <- 20
		ch3 <- 30
	}()
	for sum1 := range ch4 {
		println(sum1)
	}

}
func variadichan(chs ...chan int) chan int {
	ch1 := make(chan int)
	go func() {
		sum := 0
		for _, ch := range chs {
			sum += <-ch

		}
		ch1 <- sum

		close(ch1)

	}()
	return ch1
}
