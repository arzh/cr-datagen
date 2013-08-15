package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

func splitString(s string, n int) []string {
	strings := make([]string, 0)

	l := len(s)
	start, end := 0, n-1
	for {
		if end >= l {
			end = l
		}

		if start == l && end == l {
			break
		}

		//fmt.Println("Start:", start, "End:", end)
		strings = append(strings, s[start:end])

		if end == l {
			break
		}
		start += n
		end += n
	}
	return strings
}

func printFmtData(fi io.Writer, name string, data []string) {
	fmt.Fprintln(fi, name, `:= [...]string {`)

	l := len(data)
	for i, e := range data {
		fmt.Fprintf(fi, "`%s`", e)
		if i < l-1 {
			fmt.Fprintln(fi, ",")
		}
	}

	fmt.Fprintln(fi, "}")
}

func main() {
	fmt.Println("Generating string, 80000 characters")
	s := ""
	for i := 0; i < 80000; i++ {
		r := rand.Int31n(93) + 33
		if r != 96 {
			s += string(r)
		} else {
			i -= 1
		}
	}
	fmt.Println("String generated")

	if len(s) != 80000 {
		fmt.Println("String is short")
	}

	fmt.Println("Splitting strings, 80")
	rawData := splitString(s, 80)

	fmt.Println("Splitting strings,  125")
	testData := splitString(s, 125)

	fmt.Println("Writting file")

	fi, err := os.Create("cspTestData.go")
	if err != nil {
		fmt.Println("Shits fucked")
	} else {

		fmt.Fprintln(fi, "package main\n")

		fmt.Fprint(fi, "func csp_getTestData() ([]string, []string) {\n\n")

		printFmtData(fi, "data", rawData)

		fmt.Fprint(fi, "\n\n")

		printFmtData(fi, "testData", testData)

		fmt.Fprint(fi, "\n\nreturn data, testData\n\n}")
	}

	fi.Close()
}
