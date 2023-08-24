package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/gavsidhu/templ8/pkg"
)

var dirFlag string
var nameFlag string

func NewAddCommand() *pkg.Command {
	cmd := &pkg.Command{
		Flags:   flag.NewFlagSet("add", flag.ExitOnError),
		Execute: addFunc,
	}

	cmd.Flags.StringVar(&dirFlag, "dir", "", "Add a directory as a template")
	cmd.Flags.StringVar(&nameFlag, "name", "", "Specify a name for the template")

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, addUsage)
	}

	return cmd
}

var addFunc = func(cmd *pkg.Command, args []string) {
	if nameFlag != "" {
		if dirFlag != "" {
			fmt.Println("Directory:", dirFlag)
		}
		fmt.Println("Name:", nameFlag)
	}
}

var addUsage = `Add a template from a file path or URL.

Usage: brief add [OPTIONS] TEMPLATE

Options:
	--file	path to an existing template file
`
