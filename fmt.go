package format

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile("%<([a-zA-Z0-9_]+)>[.0-9]*[xsvTtbcdoqXxUeEfFgGp]")

// Printf support named format
func Printf(format string, params map[string]interface{}) {
	f, p := parse(format, params)
	fmt.Printf(f, p...)
}

// Printfln support named format
func Printfln(format string, params map[string]interface{}) {
	f, p := parse(format, params)
	fmt.Printf(f, p...)
	fmt.Println()
}

// Sprintf support named format
func Sprintf(format string, params map[string]interface{}) string {
	f, p := parse(format, params)
	return fmt.Sprintf(f, p...)
}

// Sprintfln support named format
func Sprintfln(format string, params map[string]interface{}) string {
	f, p := parse(format, params)
	return fmt.Sprintf(f+"\n", p...)
}

func parse(format string, params map[string]interface{}) (string, []interface{}) {
	f, n := reformat(format)
	var p []interface{}
	for _, v := range n {
		p = append(p, params[v])
	}
	return f, p
}

func reformat(f string) (string, []string) {
	i := re.FindAllStringSubmatchIndex(f, -1)

	ord := []string{}
	pair := []int{0}
	for _, v := range i {
		ord = append(ord, f[v[2]:v[3]])
		pair = append(pair, v[2]-1)
		pair = append(pair, v[3]+1)
	}
	pair = append(pair, len(f))
	plen := len(pair)

	out := ""
	for n := 0; n < plen; n += 2 {
		out += f[pair[n]:pair[n+1]]
	}

	return out, ord
}
