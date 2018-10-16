package main

import (
	"fmt"
	"unicode/utf8"
)

func IsPalindromeAsc(str string) bool {
	strLen := len(str)
	if strLen <= 1 {
		return true
	}
	for index := 0; index <= strLen/2; index++ {
		if str[index] != str[strLen-index-1] {
			return false
		}
	}
	return true
}

func IsPalindromeUtf(str string) bool {
	strLen := utf8.RuneCountInString(str)
	if strLen <= 1 {
		return true
	}
	for strLen > 1 {
		first, firstSize := utf8.DecodeRuneInString(str)
		last, lastSize := utf8.DecodeLastRuneInString(str)
		if first != last {
			return false
		}
		str = str[firstSize : strLen-lastSize] //here,  pay attention, slice cut will subtract 1
		strLen = utf8.RuneCountInString(str)
	}
	return true
}

func main() {
	str := "asdfghjkjhgfdsa"
	fmt.Printf("The str is %s, palin status %v \n", str, IsPalindromeAsc(str))
	fmt.Printf("The str is %s, palin status %v \n", str, IsPalindromeUtf(str))

}
