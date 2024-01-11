package task

import (
	"strconv"
	"unicode"
)

func konv(a, b string) int64 {

	arr_a := []rune(a)
	arr_b := []rune(b)
	for i := 0; i < len(arr_a); i++ {
		if !unicode.IsDigit(arr_a[i]) && i < len(arr_a) {
			arr_a = append(arr_a[:i], arr_a[i+1:]...)
			i--
		} else if !unicode.IsDigit(arr_a[i]) && i == len(arr_a) {
			arr_a = arr_a[:len(arr_a)-1]
			i--
		}
	}
	for i := 0; i < len(arr_b); i++ {
		if !unicode.IsDigit(arr_b[i]) && i < len(arr_b) {
			arr_b = append(arr_b[:i], arr_b[i+1:]...)
			i--
		} else if !unicode.IsDigit(arr_b[i]) && i == len(arr_b) {
			arr_b = arr_b[:len(arr_b)-1]
			i--
		}
	}
	arr_a_int, _ := strconv.Atoi(string(arr_a))
	arr_b_int, _ := strconv.Atoi(string(arr_b))

	c := int64(arr_a_int + arr_b_int)
	return c

}
