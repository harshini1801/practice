package main
func main(){


}
type iarea interface{
	Area() float32
}
type iperimeter interface{
	perimeter() float32
}
func (iarea1 iarea){
	a:=iarea1.Area()
	fmt.Println(a)
}
func perimeter(ipermeter1 iperimeter){
	p:=iperimeter1.perimeter()
	fmt.Println(p)
}