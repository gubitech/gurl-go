package main

import (
	"fmt"

	curl "github.com/andelf/go-curl"
)

// make a callback function
func output(buf []byte, userdata interface{}) bool {
  println(string(buf))
  return true
}

func execute(verb string, version string, server string) {
  easy := curl.EasyInit()
  defer easy.Cleanup()

  easy.Setopt(curl.OPT_USERAGENT, version)
  easy.Setopt(curl.OPT_URL, *flagServer)

  easy.Setopt(curl.OPT_WRITEFUNCTION, output)

  if err := easy.Perform(); err != nil {
    fmt.Printf("ERROR: %v\n", err)
  }
}
