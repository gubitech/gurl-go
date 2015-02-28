package main

import (
	"fmt"
	"os"
	"flag"

	curl "github.com/andelf/go-curl"
	"github.com/spf13/cobra"
	"github.com/rakyll/globalconf"
)

// this is set by the build script
var release string

type flagValue struct {
	str string
}

func (f *flagValue) String() string {
	return f.str
}

func (f *flagValue) Set(value string) error {
	f.str = value
	return nil
}

func newFlagValue(val string) *flagValue {
	return &flagValue{str: val}
}

var (
	flagServer = flag.String("server", "", "Current server to execute calls against.")
  flagAddress = flag.String("addr", "", "Address of the person.")
)

func main() {
	conf, err := globalconf.New("gurl")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	conf.ParseAll()

	version := fmt.Sprintf("Gurl: v%s", release)

	GurlCmd := &cobra.Command{
		Use:   "gurl",
		Short: "Gurl is a wrapper around cURL",
		Long:  "blah blah blah",
		Run: func(cmd *cobra.Command, args []string) {
			if *flagServer == "" {
				println("ERROR: server not set! Call the `set` command first.")
				fmt.Printf("%s", *flagAddress)
				os.Exit(-1)
			} else {
				fmt.Printf("Using %s\n", *flagServer)
			}

			easy := curl.EasyInit()
			defer easy.Cleanup()

			easy.Setopt(curl.OPT_USERAGENT, version)
			easy.Setopt(curl.OPT_URL, *flagServer)

			// make a callback function
			fooTest := func(buf []byte, userdata interface{}) bool {
				println("DEBUG: size=>", len(buf))
				println("DEBUG: content=>", string(buf))
				return true
			}

			easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

			if err := easy.Perform(); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		},
	}

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Set the server to use",
		Long:  "This server is what all your requests are made to",
		Run: func(cmd *cobra.Command, args []string) {
			conf.Set("", &flag.Flag{Name: "server", Value: newFlagValue(args[0])})
			fmt.Printf("Set <%s>\n", args[0])
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of gurl",
		Long:  "The version number of gurl also includes the User-Agent that's used for requests",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Current server: %s\n", *flagServer)
			fmt.Printf("User-Agent: %s\n", version)
		},
	}

	GurlCmd.AddCommand(setCmd)
	GurlCmd.AddCommand(versionCmd)

	GurlCmd.Execute()
}
