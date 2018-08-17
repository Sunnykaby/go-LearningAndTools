package main

import (
	"crypto/sha256"
	"fmt"
	"math"
	"regexp"
)

func generateBitsWithByte(src []byte) (bits [256]int) {
	len_max := 7
	len, _bit := 0, 0
	for index, value := range src {
		// fmt.Printf("Index %d, Byte %X Unicode %U Int %d \n", index, value, value, value)
		for len = 0; len <= len_max; len++ {
			// fmt.Printf("%U \n", value)
			if value&0x0001 == 1 {
				_bit = 1
			} else {
				_bit = 0
			}
			bits[index*8+len_max-len] = _bit
			value >>= 1
		}
	}
	return bits
}

func groupToInt(bits [256]int, len int) []int {
	groupNum := int(math.Floor((256) / float64(len)))
	groups := make([]int, groupNum)
	groupValue := 0

	for index := 0; index < groupNum; index++ {
		for _index := 0; _index < len && (index*len+_index) < 256; _index++ {
			// fmt.Printf("Index %d bits %d bitValue %d \n", _index, uint(len-_index-1), 1<<uint(len-_index-1))
			if bits[index*len+_index] == 1 {
				groupValue += 1 << uint(len-_index-1)
			}
		}
		groups[index] = groupValue
		groupValue = 0
	}
	return groups
}

func getN(num int) int {
	n := 0
	for sum := 1; sum < num; sum <<= 1 {
		// fmt.Printf("Sum %d, num %d \n", sum, num)
		n++
	}
	return n
}

func generateBitsWithString(src string) (bitStr string) {
	if matched, err := regexp.MatchString("[0|1|2|3|4|5|6|7|8|9|a|b|c|d|e|f]*", src); err != nil && matched == false {
		fmt.Println("Invalid Input")
		return bitStr
	}

	for index, ch := src {
		
	}
}

func main() {
	h := sha256.New()
	h.Write([]byte("this is a test1"))
	h256 := h.Sum(nil)
	fmt.Printf("%x \n", h256)
	// var bits []int
	bits := generateBitsWithByte(h256)
	for index, value := range bits {
		fmt.Printf("Index %d, Value %d \n", index, value)
	}
	// fmt.Printf("%c %X %d", '0', '0', '0')
	//Two method
	num := 7
	len := getN(num)
	groups := groupToInt(bits, len)

	for index, value := range groups {
		fmt.Printf("Index %d, Data %d \n", index, value)
	}

}
