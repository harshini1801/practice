package main

import "fmt"

func main() {
	fn1 := "victoria"
	fn2 := "secrets"
	str1 := getfulname(&fn1, &fn2)
	str2 := getfulname(&fn2, nil)
	fmt.Println(*str1, *str2)

}
func getfulname(firstname, lastname *string) *string {
	if firstname == nil || *firstname == " " {
		panic("firstname cannot br nil")
	}
	if lastname == nil || *lastname == "" {
		panic("last name cannot be nil")
	}
	str := *firstname + *lastname
	return &str
}

//when to raise panic
//without a succesful execution,if your application cannot proced further then you can consider rasising panic
