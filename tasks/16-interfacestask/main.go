package main

import "fmt"

func main() {
	it1 := mymap{m: map[string]string{"harshi": "girl", "dileep": "boy"}}
	var intf interface1
	intf = &it1
	k1 := intf.tostring()
	fmt.Println(k1)
	k2 := intf.count()
	fmt.Println(k2)

}

type mymap struct {
	m map[string]string
	/*
		result1 string
		result2 any
	*/
}

/*
	type arr struct {
		array  []string
		result string
	}
*/
type interface1 interface {
	tostring() string
	count() int
}

/*
	type interface2 interface {
		count() int
	}
*/
func (m1 *mymap) tostring() string {
	str := ""
	for index, value := range m1.m {
		k := fmt.Sprintf("key :%s,value:%s", index, value)
		str = str + " " + k
	}
	return str

}
func (m1 *mymap) count() int {
	result := len(m1.m)
	return result
}
