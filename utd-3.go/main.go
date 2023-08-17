package main

import "fmt"

func main() {
	num1:=100
	var num2 myint=200
	var num3 myint2=300
	str1:=num2.ToString()
	str2:=myint(num1).ToString()
	str3:=myint(num3).ToString()



}

type myint int     //a new type which is user defined data type
type integer = int //just an alias for int
type myint2 myint
func (mi myint) ToString() string {
	return fmt.Sprintf(mi)
}
type square float32
func(s square)area()float32{
	return float32(s*s)
}
