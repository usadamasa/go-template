package log

import (
	"bytes"
	"testing"
)

func swapWriters(t *testing.T) (out, errBuf *bytes.Buffer) {
	t.Helper()
	out, errBuf = &bytes.Buffer{}, &bytes.Buffer{}
	origOut, origErr := Out, Err
	Out, Err = out, errBuf
	t.Cleanup(func() { Out, Err = origOut, origErr })
	return out, errBuf
}

func TestStdoutFunctions(t *testing.T) {
	tests := []struct {
		name string
		call func()
		want string
	}{
		{name: "Println", call: func() { Println("hello", 1) }, want: "hello 1\n"},
		{name: "Printf", call: func() { Printf("value: %d", 42) }, want: "value: 42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, errBuf := swapWriters(t)
			tt.call()
			if got := out.String(); got != tt.want {
				t.Errorf("stdout = %q, want %q", got, tt.want)
			}
			if errBuf.Len() != 0 {
				t.Errorf("stderr = %q, want empty", errBuf.String())
			}
		})
	}
}

func TestStderrFunctions(t *testing.T) {
	tests := []struct {
		name string
		call func()
		want string
	}{
		{name: "Errorln", call: func() { Errorln("oops", 2) }, want: "oops 2\n"},
		{name: "Errorf", call: func() { Errorf("failed: %s", "x") }, want: "failed: x"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, errBuf := swapWriters(t)
			tt.call()
			if got := errBuf.String(); got != tt.want {
				t.Errorf("stderr = %q, want %q", got, tt.want)
			}
			if out.Len() != 0 {
				t.Errorf("stdout = %q, want empty", out.String())
			}
		})
	}
}
