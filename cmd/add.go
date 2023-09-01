package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gavsidhu/templ8/helpers"
	"github.com/gavsidhu/templ8/pkg"
)

func NewAddCommand() *pkg.Command {
	cmd := &pkg.Command{
		Flags:   flag.NewFlagSet("add", flag.ExitOnError),
		Execute: addFunc,
	}

	cmd.Flags.String("dir", "", "Add a directory as a template")
	cmd.Flags.String("file", "", "Add a file as a template")
	cmd.Flags.String("name", "", "Specify a name for the template")

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, addUsage)
	}

	return cmd
}

var addFunc = func(cmd *pkg.Command, args []string) {
	var nameFlag = cmd.Flags.Lookup("name").Value.String()
	var dirFlag = cmd.Flags.Lookup("dir").Value.String()
	var fileFlag = cmd.Flags.Lookup("file").Value.String()

	if nameFlag != "" {
		if dirFlag != "" && fileFlag != "" {
			fmt.Fprintln(os.Stderr, "Error: Cannot use both -dir and -file flags together.")
			cmd.Flags.Usage()
			os.Exit(1)
		}

		if dirFlag != "" {
			var isValid, err = helpers.ValidateDirectory(dirFlag)

			if err != nil {
				helpers.ErrAndExit(err)
			}

			if !isValid {
				helpers.ErrAndExit("Error: Invalid directory path")
			}

			destinationDir := filepath.Join("templates/dir", *&nameFlag)

			if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
				helpers.ErrAndExit(err)
			}

			if err := helpers.CopyDirectory(dirFlag, destinationDir); err != nil {
				helpers.ErrAndExit(err)
			}

			fmt.Printf("Added directory '%s' to '%s'\n", dirFlag, destinationDir)

		} else if fileFlag != "" {
			fmt.Println(fileFlag)
			isValid, err := helpers.ValidateFile(fileFlag)
			if err != nil {
				helpers.ErrAndExit(err)
			}
			if !isValid {
				helpers.ErrAndExit(err)
			}

			originalFileName := filepath.Base(fileFlag)

			destinationDir := "templates/files/"
			if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
				fmt.Println("Error:", err)
				return
			}

			destinationFile := filepath.Join(destinationDir, nameFlag+filepath.Ext(fileFlag))
			if err := helpers.CopyFile(fileFlag, destinationFile); err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Printf("Added file '%s' as '%s' under templates/files/\n", fileFlag, originalFileName)

		}
	}
}

var addUsage = `Add a directory or file as a template.

Usage: templ8 add (--dir <directory_path> | --file <file_path>) --name <template_name>

Example: templ8 add --file example.md --name example-template

Options:
	--file	path to an existing file
	--dir	path to an existing directory
	--name	Specify a name for the template
`
