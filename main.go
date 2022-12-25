package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()

	text := strings.TrimSpace(reader.Text())

	if len(text) == 0 {
		fmt.Println(0)

	} else {

		numOfWords := strings.Split(text, " ")

		fmt.Println(len(numOfWords))
	}

}