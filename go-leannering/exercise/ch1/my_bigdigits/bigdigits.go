// Copyright © 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func usage() {
	fmt.Printf("usage: %s [-b|--bar] <whole-number>\n-b --bar draw an underbar and an overbar\n",
		filepath.Base(os.Args[0]))
	os.Exit(1)
}

func printBar(digitStr string) {
	lenth := 0
	for index := range digitStr {
		lenth += len(bigDigits[digitStr[index]-'0'][0]) + 2
	}
	line := strings.Repeat("*", lenth-2)
	fmt.Println(line)
}

func main() {
	isBar := false
	stringOfDigits := ""
	switch len(os.Args) {
	case 1:
		usage()
	case 2:
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			usage()
		} else {
			if matched, err := regexp.MatchString("[0-9]*", os.Args[1]); err == nil && matched == true {
				stringOfDigits = os.Args[1]
			} else {
				usage()
			}
		}
	case 3:
		if matched, err := regexp.MatchString("[0-9]*", os.Args[2]); err == nil && matched == true {
			stringOfDigits = os.Args[2]
		} else {
			usage()
		}
		if os.Args[1] == "-b" || os.Args[1] == "--bar" {
			isBar = true
		}
	}
	if isBar {
		printBar(stringOfDigits)
	}
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
	if isBar {
		printBar(stringOfDigits)
	}
}

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
		"   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}
