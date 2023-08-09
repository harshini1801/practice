package main

func main() {
	//num := 32
	num := 28
	switch {
	case num%2 == 0:
		println("div by 2")
		fallthrough
	case num%4 == 0:
		println("div by 4")
		fallthrough
	case num%8 == 0:
		println("div by 8")
	}

}
