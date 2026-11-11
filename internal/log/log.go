// Package log provides console output helpers for the CLI.
// Use this package instead of fmt.Print* so that forbidigo can keep stray output out of the codebase.
//
// Only functions that something calls are exported; add one together with its first caller.
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

// Errorln prints to Err with a newline.
func Errorln(a ...any) {
	_, _ = fmt.Fprintln(Err, a...)
}
