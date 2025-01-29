package logger

import (
	"bytes"
	"testing"
)

func TestBuf(t *testing.T) {
	buf := bytes.Buffer{}
	buf.WriteString("hello")
	buf.WriteString("world")
	data := buf.Bytes()
	t.Log(string(data))
	buf.WriteString("test")
	t.Log(string(data))
	buf.Reset()
	buf.WriteString("test")
	t.Log(string(data))

}
