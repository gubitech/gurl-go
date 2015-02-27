package main

import (
    "fmt"
    curl "github.com/andelf/go-curl"
    "github.com/spf13/cobra"
)

var release string

func main() {
  version := fmt.Sprintf("Gurl: v%s", release)
  GurlCmd := &cobra.Command{
    Use:   "gurl",
    Short: "Gurl is a wrapper around cURL",
    Long: "blah blah blah",
    Run: func(cmd *cobra.Command, args []string) {
      easy := curl.EasyInit()
      defer easy.Cleanup()

      easy.Setopt(curl.OPT_USERAGENT, fmt.Sprintf("Gurl: v %s", release))
      easy.Setopt(curl.OPT_URL, "https://api.github.com/")

      // make a callback function
      fooTest := func (buf []byte, userdata interface{}) bool {
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

  versionCmd := &cobra.Command {
      Use:   "version",
      Short: "Print the version number of gurl",
      Long:  "The version number of gurl is also the User-Agent that's used for requests",
      Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(version)
      },
  }

  GurlCmd.AddCommand(versionCmd)

  GurlCmd.Execute()
}
