package main

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func sum(first_num_str string, second_num_str string) string {
	result_string := ""
	f_len := len(first_num_str)
	s_len := len(second_num_str)
	//max(f_len, s_len)
	iter_len := f_len
	if f_len < s_len {
		iter_len = len(second_num_str)
	}
	remainder := 0
	// Code smells: Trilogy
	for i := 0; i < iter_len; i++ {
		// length hackery: code smell the epilogue
		f_digit := 0
		s_digit := 0
		var f_err error
		var s_err error

		// why.
		f_iter := f_len - 1 - i
		if f_iter < 0 {
			f_digit = 0
		} else {
			// string() is necessary because indexing gives the ascii number of that char... lol
			// Atoi is technically - '0' (see Atoi in C documentation or _Atoi64 in C# documentation)
			f_digit, f_err = strconv.Atoi(string(first_num_str[f_iter]))
		}

		s_iter := s_len - 1 - i
		if s_iter < 0 {
			s_digit = 0
		} else {
			// Same note as above
			s_digit, s_err = strconv.Atoi(string(second_num_str[s_iter]))
		}

		if f_err != nil || s_err != nil {
			panic("Error while parsing first number's or second number's digit")
		}

		// massive code smells
		overflow := false

		digit := f_digit + s_digit + remainder
		remainder = 0
		// remainder stuff
		if digit > 9 {
			digit %= 10
			remainder = 1
			if s_iter == 0 && f_iter == 0 {
				overflow = true
			}
		}
		fmt.Println("Kết quả hàng thứ", i+1, "là:", digit)

		// massive code smells pt.2: the overflowing of the digits
		if overflow {
			result_string = fmt.Sprint(digit) + result_string
			result_string = fmt.Sprint(1) + result_string
			fmt.Println("Kết quả hàng thứ", i+2, "là:", 1)
		} else {
			result_string = fmt.Sprint(digit) + result_string
		}

	}
	return result_string
}

func main() {
	first_num := "192780"
	second_num := "33221"
	log.SetLevel(log.DebugLevel)
	fmt.Print(sum(first_num, second_num))

}
