package main

const pi float32 = 3.14

// usually constants are stored in data segment
// only at compile time the value to the const has to be declared
const sq1pi = pi * pi

var pi1 float32 = 3.14

const sqpi = pi1 * pi1 //this is invalid because pi1 is variable
// group of constants
const (
	days   = 7
	months = 7
	hours  = 24
)

func main() {
	println(sq1pi)
	println(sqpi)

}
