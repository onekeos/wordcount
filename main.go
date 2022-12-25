package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error: %d", 0)
	}
	
	numOfWords := strings.Split(text, " ")
	fmt.Println(len(numOfWords))

}