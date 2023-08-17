package main

import "fmt"

func main() {
	fn1 := "victoria"
	fn2 := "secrets"
	defer recoverme()
	str2 := getfulname(&fn2, nil)
	str1 := getfulname(&fn1, &fn2)
	fmt.Println(*str1)
	fmt.Println(*str2)

}
func getfulname(firstname, lastname *string) *string {
	defer recoverme()
	if firstname == nil || *firstname == " " {
		panic("firstname cannot br nil")
	}
	if lastname == nil || *lastname == "" {
		panic("last name cannot be nil")
	}
	str := *firstname + *lastname
	return &str
}

func recoverme() {
	if p := recover(); p != nil {
		fmt.Println("panic recovered")
	}
}
