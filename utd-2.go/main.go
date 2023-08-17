package main

import "fmt"

func main() {
	P1 := Person{id: 1, name: "harshi", email: "hponna@", Address: Address{line1: "4", city: "ananthapur", state: "AP"}}
	P1.city = "anan"
	fmt.Println(P1)

}

type Person struct {
	id    int
	name  string
	email string
	Address
}
type Address struct {
	line1, city, state string
}
