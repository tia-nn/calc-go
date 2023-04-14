package main

import (
	"bufio"
	"os"
	// "github.com/Code-Hex/dd"
)

func main() {

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {

		tokens := Tokenize(s.Text())
		// println(dd.Dump(tokens))

		node, _ := Parse(tokens)
		// println(dd.Dump(node))

		println(node.calc())

	}

}
