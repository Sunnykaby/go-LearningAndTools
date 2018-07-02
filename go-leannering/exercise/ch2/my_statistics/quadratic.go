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
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Quadratic</title>
<body><h3>Quadratic Equation Solver</h3>
<p>Solves equations of the form a<i>x</i>² + b<i>x</i> + c</p>`
	form = `<form action="/" method="POST">
<input type="text" name="a" size="1"><label for="a"><i>x</i>²</label> +
<input type="text" name="b" size="1"><label for="b"><i>x</i></label> +
<input type="text" name="c" size="1"><label for="c"> →</label>
<input type="submit" name="calculate" value="Calculate">
</form>`
	pageBottom = "</body></html>"
	errorF     = `<p class="error">%s</p>`
	solution   = "<p>%s → %s</p>"
)

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Must be called before writing response
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, errorF, err)
	} else {
		if floats, message, ok := processRequest(request); ok {
			quest := formatFormulate(request.Form)
			x1, x2 := solve(floats)
			solution := formatSolutions(x1, x2)
			fmt.Fprint(writer, quest, solution)
		} else if message != "" {
			fmt.Fprintf(writer, errorF, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([3]float64, string, bool) {
	var floats [3]float64
	count := 0
	for index, key := range []string{"a", "b", "c"} {
		if param, found := request.Form[key]; found && len(param) > 0 {
			if param[0] != "" {
				if tar, err := strconv.ParseFloat(param[0], 64); err != nil {
					return floats, param[0] + "is invalid Float", false
				} else {
					floats[index] = tar
				}
			} else {
				request.Form[key][0] = "0"
				floats[index] = 0
			}
			count++
		}
	}
	if count < 3 {
		return floats, "", false
	}
	if EqualFloat(floats[0], 0, -1) {
		return floats, "The index of x2 can't be zero", false
	}
	return floats, "", true
}

func formatFormulate(form map[string][]string) string {
	return fmt.Sprintf("%s<i>x</i>² + %s<i>x</i> + %s", form["a"][0],
		form["b"][0], form["c"][0])
}

func solve(floats [3]float64) (x1 complex128, x2 complex128) {
	a, b, c := complex(floats[0], 0), complex(floats[1], 0), complex(floats[2], 0)
	root := cmplx.Sqrt(cmplx.Pow(b, 2) - 4*a*c)
	x1 = (-b + root) / (2 * a)
	x2 = (-b - root) / (2 * a)
	return x1, x2
}

func formatSolutions(x1 complex128, x2 complex128) string {
	if EqualComplex(x1, x2) {
		return fmt.Sprintf("<i>x</i>=%f", x1)
	} else {
		return fmt.Sprintf("<i>x</i>=%f or <i>x</i>=%f", x1, x2)
	}
}

func EqualFloat(a, b, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}
	return math.Abs(a-b) <= (limit * math.Min(math.Abs(a), math.Abs(b)))
}

func EqualComplex(a, b complex128) bool {
	return EqualFloat(real(a), real(b), -1) && EqualFloat(imag(a), imag(b), -1)
}
