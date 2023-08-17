package main

import (
	"os"
	"syscall"
)

func main() {
	/*
		read1 :=open("read.txt")
		bytes, err := os.ReadFile(read1)
		ch:=make(chan int)
		ch1:=make(chan int)
		ch2:=make(chan int)
		ch3:=-make(chan int)
		ch4:=make(chan int)
		ch5:=make(chan int)
		for line:=range bytes{
			ch<-line
			for
		}
	*/
	os.OpenFile("read.txt", syscall.O_RDONLY, 0644)
	
	
	
