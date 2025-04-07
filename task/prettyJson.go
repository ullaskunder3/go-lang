package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func PrettyJSON() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <json_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	var formattedJSON map[string]interface{}
	if err := json.Unmarshal(data, &formattedJSON); err != nil {
		fmt.Printf("Invalid JSON: %v\n", err)
		os.Exit(1)
	}

	prettyJSON, err := json.MarshalIndent(formattedJSON, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(prettyJSON))
}
