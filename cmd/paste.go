package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gavsidhu/templ8/helpers"
	"github.com/gavsidhu/templ8/pkg"
)

func NewPasteCommand() *pkg.Command {
	cmd := &pkg.Command{
		Flags:   flag.NewFlagSet("paste", flag.ExitOnError),
		Execute: pasteFunc,
	}

	cmd.Flags.Bool("dir", false, "Add a directory as a template")
	cmd.Flags.Bool("file", false, "Add a file as a template")
	cmd.Flags.String("name", "", "The a name for the template")

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, pasteUsage)
	}

	return cmd
}

var pasteFunc = func(cmd *pkg.Command, args []string) {
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

			currentDirectory, err := os.Getwd()
			if err != nil {
				helpers.ErrAndExit(err)
			}

			currentDirectory = filepath.Join(currentDirectory, nameFlag)

			if err != nil {
				helpers.ErrAndExit(err)
			}

			helpers.PasteDirectory(template, currentDirectory)

			fmt.Printf("Created directory template: '%s' in '%s'\n", nameFlag, currentDirectory)

		} else if fileFlag == "true" {
			template := filepath.Join("templates/files", nameFlag)

			_, err := os.Stat(template)

			if err != nil {
				helpers.ErrAndExit("Template file not found")
			}

			currentDirectory, err := os.Getwd()

			if err != nil {
				helpers.ErrAndExit(err)
			}

			currentDirectory = filepath.Join(currentDirectory, nameFlag)

			helpers.PasteFile(template, currentDirectory)

			fmt.Printf("Created a file template: '%s' in '%s'", nameFlag, currentDirectory)
		} else {
			fmt.Fprintln(os.Stderr, "Error: Please specify either -dir or -file flag.")
			cmd.Flags.Usage()
			os.Exit(1)
		}
	}
}

var pasteUsage = `Paste a template in the current directory.

Usage: templ8 paste (--dir=true | --file=true) --name <template_name>

Example: templ8 paste --file=true --name example.md

Options:
	--file	whether a template is a file. (default: false)
	--dir	whether a template is a directory. (default: false)
	--name	The name of the template you want to paste
`
