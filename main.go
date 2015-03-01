package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rakyll/globalconf"
	"github.com/spf13/cobra"
)

var (
	release    string // this is set by the build script
	flagServer = flag.String("server", "", "Current server to execute calls against.")
)

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

func checkServer(server string) {
	if server == "" {
		ErrorPrinter("ERROR: server not set! Call the `set` command first.")
		os.Exit(-1)
	}
}

func main() {
	conf, err := globalconf.New("gurl")
	if err != nil {
		ErrorPrinter(err)
	}
	conf.ParseAll()

	version := fmt.Sprintf("Gurl: v%s", release)

	GurlCmd := &cobra.Command{
		Use:   "gurl",
		Short: "Gurl is a wrapper around cURL",
		Long:  "Gurl is a wrapper around cURL",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)

			execute("GET", version, *flagServer)
		},
	}

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Performs a GET request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)

			execute("GET", version, *flagServer)
		},
	}

	headCmd := &cobra.Command{
		Use:   "head",
		Short: "Performs a HEAD request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)

			execute("HEAD", version, *flagServer)
		},
	}

	optionsCmd := &cobra.Command{
		Use:   "options",
		Short: "Performs a OPTIONS request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)

			execute("OPTIONS", version, *flagServer)
		},
	}

	postCmd := &cobra.Command{
		Use:   "post",
		Short: "Performs a POST request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)

			execute("POST", version, *flagServer)
		},
	}

	patchCmd := &cobra.Command{
		Use:   "post",
		Short: "Performs a PATCH request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)

			execute("PATCH", version, *flagServer)
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Performs a DELETE request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)

			execute("DELETE", version, *flagServer)
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

	GurlCmd.AddCommand(getCmd)
	GurlCmd.AddCommand(headCmd)
	GurlCmd.AddCommand(optionsCmd)
	GurlCmd.AddCommand(postCmd)
	GurlCmd.AddCommand(patchCmd)
	GurlCmd.AddCommand(deleteCmd)
	GurlCmd.AddCommand(setCmd)
	GurlCmd.AddCommand(versionCmd)

	GurlCmd.Execute()
}
