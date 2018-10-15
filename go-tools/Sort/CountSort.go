package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a count array, store the total count
func genCountArray(src []int) (res []int, min int, max int) {
	max, min = findMaxMin(src)
	//To avoid the start num is not 1
	res = make([]int, max-min+1)

	for _, value := range src {
		res[value-min]++
	}
	return res, min, max
}

func findMaxMin(src []int) (max int, min int) {
	if len(src) == 0 {
		return -1, -1
	}
	min = src[0]
	for _, value := range src {
		if value > max {
			max = value
		} else if value < min {
			min = value
		}
	}
	return
}

func genRandomArray(lenArr int, max int, adding int) []int {
	rand.Seed(time.Now().Unix())
	res := make([]int, lenArr)

	for i := 0; i < lenArr; i++ {
		res[i] = rand.Intn(max) + adding
	}
	return res
}

func main() {
	input := genRandomArray(50, 30, 10)
	fmt.Println(input)
	countArr, min, max := genCountArray(input)
	fmt.Printf("The count array: \n%v\nThe min is : %d\nThe max is : %d\n", countArr, min, max)
	//Solution 1, scan the count array, output the sorted array, not stable
	i := 0
	for index, value := range countArr {
		for value > 0 {
			input[i] = index + min
			value--
			i++
		}
	}
	fmt.Printf("The sorted array is : \n%v\n", input)
	//Solution 2, scan the origin array, output the sorted array , stable
	res := make([]int, len(input))
	for index := 1; index < len(countArr); index++ {
		countArr[index] += countArr[index-1]
	}
	for index := len(input) - 1; index >= 0; index-- {
		res[countArr[input[index]-min]-1] = input[index]
		countArr[input[index]-min]--
	}
	fmt.Printf("The sorted array is : \n%v\n", res)
}
