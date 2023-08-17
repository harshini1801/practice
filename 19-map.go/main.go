package main

import "fmt"

// key should be unique
func main() {
	var mymap map[string]any
	mymap = make(map[string]any)
	mymap["522001"] = "gun1"
	mymap["522002"] = "gun2"
	mymap["560086"] = "gun3"
	for key, value := range mymap {
		fmt.Println("Key:", key, "Value:", value)

	}
	_, ok := mymap["522001"]
	if ok {
		fmt.Println("key exists")

	} else {
		fmt.Println("key doesnot")
	}
	delete(mymap, "522001")
	fmt.Println(mymap)

}
