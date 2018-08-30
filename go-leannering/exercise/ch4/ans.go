package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	slice_1 := []int{
		1, 3, 4, 5, 6, 2, 34, 6, 7, 34, 5, 4, 3, 6, 3, 4, 64, 5675, 74, 63, 234, 134, 2, 53, 6, 5, 354, 2, 4, 2, 4}
	fmt.Println("Src int slice : ", slice_1)
	fmt.Println("Remove the dumplicated item int slice : ", UniqueInts(slice_1))

	slice_2 := [][]int{
		{1, 2, 42, 3}, {2, 3, 4, 2}, {23, 2345, 234},
		{234, 234, 2, 2, 4, 235}, {123, 234, 2, 14, 5}, {12, 23, 345, 4, 1, 3, 7}}
	fmt.Printf("The slice_2 len is %d, data is : \n %v \n", len(slice_2), slice_2)
	fmt.Printf("Flaten the slice, data is  %v \n", Flatten(slice_2))

	slice_3 := Flatten(slice_2)
	fmt.Printf("The sclice_3 src data: %v \n", slice_3)
	fmt.Printf("The 2D with cutsize (%d) : %v \n", 4, Make2D(slice_3, 4))

	iniData := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}

	fmt.Printf("Src ini data is %v \n", iniData)
	iniMap := ParseIni(iniData)
	fmt.Printf("Parsed ini data is : %v \n", iniMap)
	PrintIni(iniMap)
}

func UniqueInts(src []int) []int {
	//Utilize the map to store the existing int
	exiInts := map[int]bool{} // == make(map[int]bool)
	results := make([]int, 0) //we didn't know the len == []int{}

	for _, val := range src {
		if _, found := exiInts[val]; found {
			continue
		}
		exiInts[val] = true
		results = append(results, val)
	}
	return results
}

func Flatten(src [][]int) []int {
	// init as a default size
	flatInts := make([]int, 0, len(src)+len(src[0]))
	for _, rowItem := range src {
		for _, colItem := range rowItem {
			flatInts = append(flatInts, colItem)
		}
	}
	return flatInts
}

func Make2D(src []int, cutSize int) [][]int {
	row := cutRow(len(src), cutSize)
	result := make([][]int, row)
	for i := 0; i < row; i++ {
		col := make([]int, cutSize)
		for j := 0; j < cutSize && i*cutSize+j < len(src); j++ {
			col[j] = src[i*cutSize+j]
		}
		result[i] = col
	}
	return result
}

func cutRow(total int, cutSize int) int {
	row := total / cutSize
	if total%cutSize > 0 {
		row++
	}
	return row
}

func ParseIni(iniData []string) map[string]map[string]string {
	result := map[string]map[string]string{}
	group := ""
	for _, str := range iniData {
		str = strings.TrimSpace(str)
		if len(str) == 0 || strings.HasPrefix(str, ";") {
			continue
		}
		if strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]") {
			group = str[1 : len(str)-1]
			groupMap := map[string]string{}
			result[group] = groupMap
		}
		if strings.Contains(str, "=") {
			pairs := strings.Split(str, "=")
			fmt.Println(pairs)
			result[group][pairs[0]] = pairs[1]
		}
	}
	return result
}

func PrintIni(iniMap map[string]map[string]string) {
	group := []string{}
	for gKey := range iniMap {
		group = append(group, gKey)
	}
	sort.Strings(group)
	for _, gKey := range group {
		subKeys := []string{}
		for key := range iniMap[gKey] {
			subKeys = append(subKeys, key)
		}
		sort.Strings(subKeys)
		fmt.Printf("[%s]\n", gKey)
		for _, key := range subKeys {
			fmt.Printf("%s=%v\n", key, iniMap[gKey][key])
		}
	}
}
