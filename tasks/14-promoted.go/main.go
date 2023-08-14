package main

import "fmt"

type Company struct {
	id, name, Website, Telehone, Fax string
	Address
}
type Address struct {
	city, line1, Street, State string
	Map
}
type Map struct {
	Lan, Lat string
}

func main() {
	c1 := Company{id: "2", name: "harshi", Website: "com", Telehone: "123", Fax: "2", Address: Address{city: "ananthapur", line1: "456", Street: "kota", Map: Map{Lan: "45deg", Lat: "64deg"}}}
	fmt.Println(c1.Address.Lan, c1.Address.Lat)
}
