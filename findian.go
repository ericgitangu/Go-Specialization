package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter a string: ")
	in := bufio.NewReader(os.Stdin)

	// Convert input to lowercase for case-insensitivity
	input, err := in.ReadString('\n')
	if err == nil {
		input = strings.ToLower(input)
		input = strings.Join(strings.Fields(input), "")
	}

	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
