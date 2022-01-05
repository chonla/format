package main

import (
	"github.com/chonla/format"
)

func main() {
	var params = map[string]interface{}{
		"sister":  "Susan",
		"brother": "Louis",
	}
	format.Printf("%<brother>s loves %<sister>s.", params)
}
