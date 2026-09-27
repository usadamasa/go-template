// Package log provides console output helpers for the CLI.
// Use this package instead of fmt.Print* so that forbidigo can keep stray output out of the codebase.
//
// The API is intentionally small: two functions for stdout and two for stderr.
// Tests swap Out / Err for a buffer instead of a logger type.
package log

import (
	"fmt"
	"io"
	"os"
)

// Out receives normal output. Err receives error output.
var (
	Out io.Writer = os.Stdout
	Err io.Writer = os.Stderr
)

// Println prints to Out with a newline.
func Println(a ...any) {
	_, _ = fmt.Fprintln(Out, a...)
}

// Printf prints formatted output to Out.
func Printf(format string, a ...any) {
	_, _ = fmt.Fprintf(Out, format, a...)
}

// Errorln prints to Err with a newline.
func Errorln(a ...any) {
	_, _ = fmt.Fprintln(Err, a...)
}

// Errorf prints formatted output to Err.
func Errorf(format string, a ...any) {
	_, _ = fmt.Fprintf(Err, format, a...)
}
