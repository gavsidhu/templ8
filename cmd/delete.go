package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gavsidhu/templ8/helpers"
	"github.com/gavsidhu/templ8/pkg"
)

func NewDeleteCommand() *pkg.Command {
	cmd := &pkg.Command{
		Flags:   flag.NewFlagSet("delete", flag.ExitOnError),
		Execute: deleteFunc,
	}

	cmd.Flags.Bool("dir", false, "Specify if the template to delete is a directory")
	cmd.Flags.Bool("file", false, "Specify if the template to delete is a file")
	cmd.Flags.String("name", "", "The name of the template to delete")

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, deleteUsage)
	}
	return cmd
}

var deleteFunc = func(cmd *pkg.Command, args []string) {
	var nameFlag = cmd.Flags.Lookup("name").Value.String()
	var dirFlag = cmd.Flags.Lookup("dir").Value.String()
	var fileFlag = cmd.Flags.Lookup("file").Value.String()

	if nameFlag != "" {
		if dirFlag == "true" && fileFlag == "true" {
			fmt.Fprintln(os.Stderr, "Error: Cannot set both -dir and -file flags as true.")
			cmd.Flags.Usage()
			os.Exit(1)
		}
		if dirFlag == "true" {
			template := filepath.Join("templates/dir", nameFlag)

			dirInfo, err := os.Stat(template)

			if err != nil || !dirInfo.IsDir() {
				helpers.ErrAndExit("Template directory not found")
			}

			err = os.Remove(template)

			if err != nil {
				helpers.ErrAndExit(err)
			}

			fmt.Printf("Deleted directory template: '%s'\n", nameFlag)
		} else if fileFlag == "true" {
			template := filepath.Join("templates/files", nameFlag)

			_, err := os.Stat(template)

			if err != nil {
				helpers.ErrAndExit(err)
			}

			err = os.Remove(template)

			if err != nil {
				helpers.ErrAndExit(err)
			}
			fmt.Printf("Deleted file template: '%s'\n", nameFlag)
		} else {
			fmt.Fprintln(os.Stderr, "Error: Please specify either -dir or -file flag.")
			cmd.Flags.Usage()
			os.Exit(1)
		}
	}
}

var deleteUsage = `Paste a template in the current directory.

Usage: templ8 delete (--dir=true | --file=true) --name <template_name>

Example: templ8 delete --file=true --name example.md

Options:
	--file	whether the template to delete is a file. (default: false)
	--dir	whether the template to delete is a directory. (default: false)
`
