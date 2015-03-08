package main

import (
	"fmt"

	curl "github.com/andelf/go-curl"
)

var (
	buffer []byte
)

func execute(verb string, version string, args []string) {
	var endpoint string

	easy := curl.EasyInit()
	defer easy.Cleanup()

	for i, a := range args {
		if i == 0 {
			endpoint = fmt.Sprint(*flagServer, a)
		} else {
			switch a {
			case "includeHeaders":
				easy.Setopt(curl.OPT_HEADERFUNCTION, func(buf []byte, userdata interface{}) bool {
					Print(buf)
					return true
				})
			case "verbose":
				easy.Setopt(curl.OPT_VERBOSE, true)
			}
		}
	}

	// always use the .netrc file
	easy.Setopt(curl.OPT_NETRC, true)

	easy.Setopt(curl.OPT_USERAGENT, version)
	easy.Setopt(curl.OPT_URL, endpoint)

	easy.Setopt(curl.OPT_WRITEFUNCTION, func(buf []byte, userdata interface{}) bool {
		buffer = append(buffer, buf...)
		return true
	})

	if err := easy.Perform(); err != nil {
		ErrorPrinter(err)
	}

	Print(buffer)
}
