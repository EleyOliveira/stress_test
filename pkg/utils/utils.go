package utils

import (
	"encoding/json"
	"fmt"
)

// FormatOutput formats the output to JSON and prints it to the console.
func FormatOutput(data interface{}) {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error formatting output:", err)
		return
	}
	fmt.Println(string(output))
}

// ParseInput reads input from the command line and unmarshals it into the provided struct.
func ParseInput(input string, output interface{}) error {
	return json.Unmarshal([]byte(input), output)
}
