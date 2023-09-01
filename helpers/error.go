package helpers

import (
	"fmt"
	"os"
)

type simpleError string

func (s simpleError) Error() string {
	return string(s)
}

func ErrAndExit(err interface{}) {
	switch err := err.(type) {
	case error:
		fmt.Fprintln(os.Stderr, "Error:", err.Error())
	case string:
		fmt.Fprintln(os.Stderr, "Error:", err)
	default:
		fmt.Fprintln(os.Stderr, "Unknown error")
	}
	os.Exit(1)
}
