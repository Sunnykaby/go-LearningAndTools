package main

import (
	"fmt"
	"os"
	"strconv"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

func main() {
	level := 0
	target := ""
	// jsonFile := ""
	switch len(os.Args) {
	case 3:
		if _level,err := strconv.Atoi(os.Args[1]); err != nil {
			level = _level
		}
		target = os.Args[2]
	default:
		fmt.Printf("Please input the right args. \n")
	}	
	fmt.Printf("The args is : length %d,  %d, %s", len(os.Args), level, target)

	//Load the json file
	path,_ := os.Getwd()
	fmt.Printf("Work Dir is '%s'", path)
	resource,err := ioutil.ReadFile(filepath.Join(path,"test.json"))
	if err != nil {
		fmt.Printf("Error %s", err)
		os.Exit(1)
	}
	var dat map[string]interface{}
	json.Unmarshal(resource, &dat)

	//这里非常坑，go的map顺序和json本身的顺序是不一致的。
	for key,_ := range dat {
		fmt.Printf("%s: \n", key)
	}
}