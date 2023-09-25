package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gavsidhu/templ8/helpers"
	"github.com/gavsidhu/templ8/pkg"
)

func NewListCommand() *pkg.Command {
	cmd := &pkg.Command{
		Flags:   flag.NewFlagSet("list", flag.ExitOnError),
		Execute: listFunc,
	}

	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listUsage)
	}

	return cmd
}

var listFunc = func(cmd *pkg.Command, args []string) {
	templateDir := filepath.Join("", "/users/shared", "templates/dir")
	fileDir := filepath.Join("", "/users/shared", "templates/files")

	dirEntries, dirErr := os.ReadDir(templateDir)
	fileEntries, fileErr := os.ReadDir(fileDir)

	if dirErr != nil || fileErr != nil {
		helpers.ErrAndExit("Error reading templates")
	}

	fmt.Println("Directory Templates:")
	if len(dirEntries) == 0 {
		fmt.Println("No directory templates")
	} else {
		for _, entry := range dirEntries {
			fmt.Printf("- %s\n", entry.Name())
		}
	}

	fmt.Println()

	fmt.Println("File Templates:")
	if len(fileEntries) == 0 {
		fmt.Println("No file templates")
	} else {
		for _, entry := range fileEntries {
			fmt.Printf("- %s\n", entry.Name())
		}
	}
}

var listUsage = `List all saved templates.

Usage: templ8 list
`
