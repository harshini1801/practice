package slice

// func reverse(s []int) []int {
// 	var rs []int
// 	for i := len(s) - 1; i <= 0; i-- {
// 		rs = append(rs, s[i])
// 	}
// 	return rs

// }
func big(s []int) int {
	big1 := s[0]
	for i := 1; i < len(s); i++ {
		if big1 < s[i] {
			big1 = s[i]
		}
	}
	return big1

}

// func small(s []int) int {
// 	small := s[0]
// 	for j := 1; j < len(s); j++ {
// 		if small > s[j] {
// 			small = s[j]
// 		}
// 	}
// 	return small

// }
// func length(s []int) int {
// 	return len(s)
// }
