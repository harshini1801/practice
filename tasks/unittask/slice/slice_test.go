package slice

import (
	"fmt"
	"testing"
)

//	func Testreverse(t *testing.T) {
//		s := []int{1, 2, 3, 4, 5, 6}
//		expected := []int{6, 5, 4, 3, 2, 1}
//		actual := reverse(s)
//		if expected == actual {
//			t.Fail()
//		}
//	}
func Testbig(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 29}
	expected := 5
	actual := big(s)
	fmt.Println(actual)
	if expected != actual {
		t.Fail()
	}

}

// func Testsmall(t *testing.T) {
// 	s := []int{1, 2, 3, 4, 5, 6}
// 	expected := 1
// 	actual := small(s)
// 	if expected != actual {
// 		t.Fail()
// 	}

// }
// func Testlength(t *testing.T) {
// 	s := []int{1, 2, 3, 4, 5, 6, 29}
// 	expected := 6
// 	actual := length(s)
// 	if expected != actual {
// 		t.Fail()
// 	}

// }
