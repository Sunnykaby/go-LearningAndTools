package main

import (
	"fmt"
	"strings"
)

func commonPrefix(src []string) string {
	lenS := len(src)
	if lenS <= 1 {
		return ""
	}
	//convert the string to []rune, make it suitable for utf8 strings
	res := []rune(src[0])
	for index := 1; index < lenS; index++ {
		tarR := []rune(src[index])
		curL := len(res)
		if curL > len(tarR) {
			curL = len(tarR)
		}
		for i := 0; i < curL; i++ {
			if res[i] != tarR[i] {
				res = res[:i]
				break
			}
		}
	}
	return string(res)
}

func commonPathPrefix(src []string) string {
	lenS := len(src)
	if lenS <= 1 {
		return ""
	}
	// sept := string(filepath.Separator)
	sept := "/"
	res := strings.Split(strings.TrimSpace(src[0]), sept)
	for index := 1; index < lenS; index++ {
		tarS := strings.Split(strings.TrimSpace(src[index]), sept)
		curL := len(tarS)
		if curL > len(res) {
			curL = len(res)
		}
		for i := 0; i < curL; i++ {
			if tarS[i] != res[i] {
				res = res[:i]
				break
			}
		}
	}
	return strings.Join(res, sept)
}

func main() {
	testData := [][]string{
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/goeg/prefix/extra"},
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/prefix/extra"},
		{"/pecan/π/goeg", "/pecan/π/goeg/prefix",
			"/pecan/π/prefix/extra"},
		{"/pecan/π/circle", "/pecan/π/circle/prefix",
			"/pecan/π/circle/prefix/extra"},
		{"/home/user/goeg", "/home/users/goeg",
			"/home/userspace/goeg"},
		{"/home/user/goeg", "/tmp/user", "/var/log"},
		{"/home/mark/goeg", "/home/user/goeg"},
		{"home/user/goeg", "/tmp/user", "/var/log"},
	}
	fmt.Printf("The src is \n%v\n, the common result is %s\n", testData[1], commonPrefix(testData[1]))
	fmt.Printf("The src is \n%v\n, the common result is %s\n", testData[1], commonPathPrefix(testData[1]))
}
