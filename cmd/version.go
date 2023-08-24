package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/gavsidhu/templ8/pkg"
)

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

var (
	build   = "???"
	version = "???"
	short   = false
)

var versionFunc = func(cmd *pkg.Command, args []string) {
	if short {
		fmt.Printf("brief version: v%s", version)
	} else {
		fmt.Printf("brief version: v%s, build: %s", version, build)
	}
	os.Exit(0)
}

var versionUsage = `Print the app version and build info for the current context.

Usage: templ8 version [options]

Options:
  --short  If true, print just the version number. Default false.
`
