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

func TestErrorln(t *testing.T) {
	out, errBuf := swapWriters(t)
	Errorln("oops", 2)
	if got, want := errBuf.String(), "oops 2\n"; got != want {
		t.Errorf("stderr = %q, want %q", got, want)
	}
	if out.Len() != 0 {
		t.Errorf("stdout = %q, want empty", out.String())
	}
}
