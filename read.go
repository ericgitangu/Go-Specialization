package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname [20] rune
	lname [20] rune
}

func main() {
	var fileName string
	fmt.Print("Enter the name of the text file: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var fname, lname [20]rune
		copy(fname[:], []rune(fields[0]))
		copy(lname[:], []rune(fields[1]))
		names = append(names, name{fname, lname})
	}

	for _, n := range names {
		fmt.Println(string(n.fname[:len(strings.Trim(string(n.fname[:]), "\x00"))]), string(n.lname[:len(strings.Trim(string(n.lname[:]), "\x00"))]))
	}
}
