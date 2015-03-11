package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rakyll/globalconf"
	"github.com/spf13/cobra"
)

var (
	release        string // this is set by the build script
	flagServer     = flag.String("server", "", "Current server to execute calls against.")
	includeHeaders bool
	verbose        bool
	header         = ""
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

func appendFlags() []string {
	var flags []string

	if includeHeaders {
		flags = append(flags, "includeHeaders")
	}

	if verbose {
		flags = append(flags, "verbose")
	}

	if len(header) > 0 {
		flags = append(flags, "Header: Custom, "+header)
	}

	return flags
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
			cmd.Usage()
		},
	}

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Performs a GET request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)
			flags := appendFlags()
			execute("GET", version, args, flags)
		},
	}

	headCmd := &cobra.Command{
		Use:   "head",
		Short: "Performs a HEAD request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)
			flags := appendFlags()
			execute("HEAD", version, args, flags)
		},
	}

	optionsCmd := &cobra.Command{
		Use:   "options",
		Short: "Performs a OPTIONS request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)
			flags := appendFlags()
			execute("OPTIONS", version, args, flags)
		},
	}

	postCmd := &cobra.Command{
		Use:   "post",
		Short: "Performs a POST request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)
			flags := appendFlags()
			execute("POST", version, args, flags)
		},
	}

	patchCmd := &cobra.Command{
		Use:   "post",
		Short: "Performs a PATCH request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)
			flags := appendFlags()
			execute("PATCH", version, args, flags)
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Performs a DELETE request",
		Run: func(cmd *cobra.Command, args []string) {
			checkServer(*flagServer)
			flags := appendFlags()
			execute("DELETE", version, args, flags)
		},
	}

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Set the server to use",
		Long:  "All your requests are made to this server",
		Run: func(cmd *cobra.Command, args []string) {
			conf.Set("", &flag.Flag{Name: "server", Value: newFlagValue(args[0])})
			fmt.Printf("Set <%s>\n", args[0])
		},
	}

	showCmd := &cobra.Command{
		Use:   "show",
		Short: "Shows the server you're using",
		Run: func(cmd *cobra.Command, args []string) {
			contents, err := ioutil.ReadFile(conf.Filename)
			fmt.Printf("Currently using: %s\n", conf.Filename)
			if err != nil {
				ErrorPrinter(err)
			}
			fmt.Print(string(contents))
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of gurl",
		Long:  "The version number of gurl also includes the User-Agent that's used for requests",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Current server: %s", *flagServer)
			fmt.Printf("User-Agent: %s", version)
		},
	}

	GurlCmd.AddCommand(getCmd)
	GurlCmd.AddCommand(headCmd)
	GurlCmd.AddCommand(optionsCmd)
	GurlCmd.AddCommand(postCmd)
	GurlCmd.AddCommand(patchCmd)
	GurlCmd.AddCommand(deleteCmd)

	GurlCmd.AddCommand(setCmd)
	GurlCmd.AddCommand(showCmd)
	GurlCmd.AddCommand(versionCmd)

	GurlCmd.PersistentFlags().BoolVarP(&includeHeaders, "include", "i", false, "Include the HTTP-header in the output")
	GurlCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Provide verbose/debug information in the output")
	GurlCmd.PersistentFlags().StringVarP(&header, "header", "H", "", "Source directory to read from")

	GurlCmd.Execute()
}
