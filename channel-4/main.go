package main

import "fmt"

func main() {
	ch := make(chan int)
	notify := make(chan struct{})
	notifysen := make(chan bool)
	count := 0
	go receiver(ch, notify, notifysen, count)
	go sender(ch, notifysen, count)
	go sender(ch, notifysen, count)
	<-notify

}
func sender(ch chan int, notifysen chan bool, count int) {
	for i := 0; i <= 100; i++ {
		closed := <-notifysen
		if !closed {
			ch <- i * i
		} else {
			if count == 1 {
				break
			} else {
				close(ch)
			}
		}
	}
}

func receiver(ch chan int, notify chan struct{}, notifysen chan bool, count int) {
	for {
		_, ok := <-ch
		if !ok {
			notifysen <- true
			count = 1
			break

		} else {
			notifysen <- false
			fmt.Println("receiver", <-ch)
		}

	}
	notify <- struct{}{}

}
