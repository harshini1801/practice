package main

import "fmt"

func main() {
	P := Person{id: 1, name: "harshi", email: "ponna@"}
	P1 := Person{}
	P1.id = 2
	P1.name = "dileep"
	P1.email = "dileep@"
	fmt.Println(P1, P)

}

type Person struct {
	id    int
	name  string
	email string
}
