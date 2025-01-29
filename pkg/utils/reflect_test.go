package utils

import (
	"testing"
)

type a struct {
	x int
}

func TestNewPointerType(t *testing.T) {
	t.Logf("%T", NewPointerType[*a]())
	t.Logf("%T", NewPointerType[*int]())
}
