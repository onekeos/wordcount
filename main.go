package main

import (
	// "bufio"
	"fmt"
	// "os"
	"strings"
	"flag"
)

func main() {

	flag.Parse()

	src := strings.Join(flag.Args(), "")

	text := strings.TrimSpace(src)

	if len(text) == 0 {
		fmt.Println(0)

	} else {

		numOfWords := strings.Split(text, " ")

		fmt.Println(len(numOfWords))
	}

}