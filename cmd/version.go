package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/gavsidhu/templ8/pkg"
)

var Build = "d22f04a64eb20e09a6fe854db7d9bce8aa8a9309"
var Version = "0.0.1"
var short = false

func NewVersionCommand() *pkg.Command {
	cmd := &pkg.Command{
		Flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: versionFunc,
	}

	cmd.Flags.BoolVar(&short, "short", false, "")
	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, versionUsage)
	}
	return cmd
}

var versionFunc = func(cmd *pkg.Command, args []string) {
	if short {
		fmt.Printf("brief version: v%s", Version)
	} else {
		fmt.Printf("brief version: v%s, build: %s", Version, Build)
	}
	os.Exit(0)
}

var versionUsage = `Print the app version and build info for the current context.

Usage: templ8 version [options]

Options:
  --short  If true, print just the version number. Default false.
`
