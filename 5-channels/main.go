package main

func main() {

}
func generatesquare(num int) chan int {
	ch := make(chan int)
	for i := 0; i <= num; i++ {
		ch <- i * i
	}
	return ch
}
