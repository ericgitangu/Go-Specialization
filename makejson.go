package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Prompt user for first name and address
	var firstName, address string
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&firstName)
	fmt.Print("Enter your address: ")
	reader := bufio.NewReader(os.Stdin)
	address, _ = reader.ReadString('\n')

	// Create a map to store the user's information
	userInfo := make(map[string]string)
	userInfo["firstName"] = firstName
	userInfo["address"] = strings.TrimSpace(address)

	// Marshal the map to JSON
	jsonData, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}
