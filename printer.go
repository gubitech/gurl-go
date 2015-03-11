package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ErrorPrinter(err interface{}) {
	fmt.Printf("ERROR: %v\n", err)
}

func isJson(buffer []byte) bool {
	buf := string(buffer)

	if string(buf[0]) == "{" {
		return true
	} else {
		return false
	}
}

func Print(buffer []byte) {
	if !isJson(buffer) {
		os.Stdout.Write(buffer)
	} else {
		var parsed map[string]interface{}

		err := json.Unmarshal(buffer, &parsed)
		if err != nil {
			ErrorPrinter(err)
		}
		b, err := json.MarshalIndent(parsed, "", "  ")
		if err != nil {
			ErrorPrinter(err)
		}
		os.Stdout.Write(b)
	}
}
