package main

import (
	curl "github.com/andelf/go-curl"
)

var buffer []byte

func execute(verb string, version string, server string) {
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_USERAGENT, version)
	easy.Setopt(curl.OPT_URL, *flagServer)

	easy.Setopt(curl.OPT_WRITEFUNCTION, func(buf []byte, userdata interface{}) bool {
		buffer = append(buffer, buf...)
		return true
	})

	if err := easy.Perform(); err != nil {
		ErrorPrinter(err)
	}

	PrintJson(buffer)
}
