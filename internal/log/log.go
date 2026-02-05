// Package log provides a simple logging interface that wraps standard library log.
// Use this package instead of fmt.Print* for console output to comply with forbidigo lint rules.
package log

import (
	"fmt"
	"io"
	"os"
)

// Logger provides formatted output to stdout/stderr.
type Logger struct {
	out io.Writer
	err io.Writer
}

// Default is the default logger writing to os.Stdout and os.Stderr.
var Default = &Logger{
	out: os.Stdout,
	err: os.Stderr,
}

// New creates a new Logger with the specified output writers.
func New(out, err io.Writer) *Logger {
	return &Logger{out: out, err: err}
}

// Print prints to stdout.
func (l *Logger) Print(a ...any) {
	fmt.Fprint(l.out, a...)
}

// Println prints to stdout with a newline.
func (l *Logger) Println(a ...any) {
	fmt.Fprintln(l.out, a...)
}

// Printf prints formatted output to stdout.
func (l *Logger) Printf(format string, a ...any) {
	fmt.Fprintf(l.out, format, a...)
}

// Error prints to stderr.
func (l *Logger) Error(a ...any) {
	fmt.Fprint(l.err, a...)
}

// Errorln prints to stderr with a newline.
func (l *Logger) Errorln(a ...any) {
	fmt.Fprintln(l.err, a...)
}

// Errorf prints formatted output to stderr.
func (l *Logger) Errorf(format string, a ...any) {
	fmt.Fprintf(l.err, format, a...)
}

// Package-level functions using the default logger.

// Print prints to stdout using the default logger.
func Print(a ...any) {
	Default.Print(a...)
}

// Println prints to stdout with a newline using the default logger.
func Println(a ...any) {
	Default.Println(a...)
}

// Printf prints formatted output to stdout using the default logger.
func Printf(format string, a ...any) {
	Default.Printf(format, a...)
}

// Error prints to stderr using the default logger.
func Error(a ...any) {
	Default.Error(a...)
}

// Errorln prints to stderr with a newline using the default logger.
func Errorln(a ...any) {
	Default.Errorln(a...)
}

// Errorf prints formatted output to stderr using the default logger.
func Errorf(format string, a ...any) {
	Default.Errorf(format, a...)
}
