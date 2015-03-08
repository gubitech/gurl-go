package main

import (
	"fmt"
	"strings"

	curl "github.com/andelf/go-curl"
)

var (
	buffer []byte
)

func execute(verb string, version string, args []string, flags []string) {
	endpoint := *flagServer + args[0]
	fmt.Printf("Hitting %s\n", endpoint)

	easy := curl.EasyInit()
	defer easy.Cleanup()

	// set the custom flags
	for _, f := range flags {
		switch f {
		case "includeHeaders":
			easy.Setopt(curl.OPT_HEADERFUNCTION, func(buf []byte, userdata interface{}) bool {
				Print(buf)
				return true
			})
		case "verbose":
			easy.Setopt(curl.OPT_VERBOSE, true)
		}
	}

	// looks like we've got some interesting data
	if len(args) > 1 {
		easy.Setopt(curl.OPT_POSTFIELDS, args[1])
		// must set
		easy.Setopt(curl.OPT_POSTFIELDSIZE, len(args[1]))
		// disable HTTP/1.1 Expect 100
		easy.Setopt(curl.OPT_HTTPHEADER, []string{"Expect:"})
	}

	// always use the .netrc file
	easy.Setopt(curl.OPT_NETRC, true)
	// always follow location
	easy.Setopt(curl.OPT_FOLLOWLOCATION, true)
	// stop after trying for a minute
	easy.Setopt(curl.OPT_TIMEOUT, 60)

	easy.Setopt(curl.OPT_USERAGENT, version)
	easy.Setopt(curl.OPT_URL, endpoint)

	easy.Setopt(curl.OPT_WRITEFUNCTION, func(buf []byte, userdata interface{}) bool {
		buffer = append(buffer, buf...)
		return true
	})

	easy.Setopt(curl.OPT_CUSTOMREQUEST, strings.ToUpper(verb))

	if err := easy.Perform(); err != nil {
		ErrorPrinter(err)
	}

	Print(buffer)
}
