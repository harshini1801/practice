package main

func some() *int {
	r := 100
	ptr := &r
	return ptr
}
func main() {
	x := some()
	*x = 300
	println(r)

}
