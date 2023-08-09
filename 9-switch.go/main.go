package main

func main() {
	day := 4
	switch day {
	case 1:
		println("sunday")
	case 2:
		println("monday")
	case 3:
		println("tuesday")
	case 4:
		println("wed")
	case 5:
		println("thurs")
	case 6:
		println("fri")
	case 7:
		println("sat")
	}
	x := 4
	println(x)
	char := 'A'
	switch char {
	case 'A', 'E', 'I', 'O', 'U':
		println(char)
	case 'a', 'e', 'i', 'o', 'u':
		println(char)
	default:
		println(char)
	}
}
