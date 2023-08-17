package main

import (
	"fmt"

	"../fileops"
)

func main() {
	filename := new(string)
	*filename = "data1.txt"
	count, err := fileops.Getfilechar(filename)
	if err != nil {
		fmt.Println("there seems to be error in handling the file")
	}
	fmt.Println(count)

}
