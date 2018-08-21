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
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Soundex Compute Test</title>
<body><h3>Soundex Compute Test</h3>
<p>Computes basic soundex for a given list of words</p>`
	form1 = `<form action="/" method="POST">
<label for="words">Words (comma or space-separated):</label><br />
<input type="text" name="words" size="50"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

var letterToDigital = []rune{0, 1, 2, 3, 0, 1, 2, 0, 0, 2, 2, 4, 5, 5, 0, 1, 2, 6, 2, 3, 0, 0, 0, 2, 0, 2}

type soundex struct {
	word     string
	soundexW string
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/test", testPage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func testPage(write http.ResponseWriter, request *http.Request) {

}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Must be called before writing response
	fmt.Fprint(writer, pageTop, form1)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if words, ok := processRequest(request); ok {
			wordsMap := getSoundex(words)
			fmt.Fprint(writer, formatSoundex(wordsMap))
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]string, bool) {
	var words []string
	if slice, found := request.Form["words"]; found && len(slice) > 0 {
		words = strings.Split(slice[0], ",")
		if len(words) == 0 {
			return words, false
		} else {
			return words, true
		}
	}
	return words, false
}

func formatSoundex(wordsMap map[string]string) string {
	var buffer bytes.Buffer
	buffer.WriteString(`<table border="1">
	<tr><th>Name</th><th>Soundex</th></tr>`)
	for key, value := range wordsMap {
		buffer.WriteString(fmt.Sprintf(`<tr><td>%s</td><td>%s</td></tr>`, key, value))
	}
	buffer.WriteString(`</table>`)

	return buffer.String()
}

func getSoundex(words []string) (wordsMap map[string]string) {
	wordsMap = make(map[string]string)
	for _, word := range words {
		wordsMap[word] = getWordSoundex(word)
	}
	return
}

func getWordSoundex(word string) (soundex string) {
	if len(word) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString(strings.ToUpper(string(word[0])))
	var preChs = make(map[rune]int8) //hash is better? maybe
	//Add first ch
	preChs[rune(word[0])] = 1
	for _, ch := range strings.ToLower(word[1:]) {
		if _, found := preChs[ch]; found {
			continue
		} else {
			fmt.Println(ch - 'a')
			buffer.WriteRune(letterToDigital[ch-'a'])
			preChs[ch] = 1
		}
	}
	//Rm the 0
	soundex = buffer.String()
	buffer.Reset()

	for _, ch := range soundex {
		if (ch - '0') == 0 {
			continue
		}
		buffer.WriteRune(ch)
	}
	//Format the code
	soundex = fmt.Sprintf("%-05s", buffer.String())
	return
}
