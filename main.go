package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gavsidhu/templ8/cmd"
	"github.com/gavsidhu/templ8/pkg"
)

var (
	Version string
	Build   string
)

func main() {

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}

	if len(os.Args) < 2 {
		usageAndExit("")
	}

	var command *pkg.Command

	switch os.Args[1] {
	case "version":
		command = cmd.NewVersionCommand()
	case "add":
		command = cmd.NewAddCommand()
	case "paste":
		command = cmd.NewPasteCommand()
	case "list":
		command = cmd.NewListCommand()
	case "delete":
		command = cmd.NewDeleteCommand()
	default:
		usageAndExit(fmt.Sprintf("templ8: '%s' is not a templ8 command.\n", os.Args[1]))
	}

	command.Init(os.Args[2:])
	command.Run()
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit((0))
}

var usage = `Usage: templ8 command [options]

A simple tool to generate and manage custom templates

Options:

Commands:
  add		Adds a template to the collection from a local file
  paste		Pastes the specified template into the current directory
  list		Lists all stored templates
  delete	Removes a stored template
  version	Prints version info to console
`
