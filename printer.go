package main

import (
	"fmt"
  "os"
  "encoding/json"
)

func ErrorPrinter(err interface{}) {
  fmt.Printf("ERROR: %v\n", err)
}

func PrintJson(buffer []byte) {
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
