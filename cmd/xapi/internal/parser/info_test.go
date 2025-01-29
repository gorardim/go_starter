package parser

import (
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestInfo(t *testing.T) {
	load, err := packages.Load(&packages.Config{
		Mode: packages.LoadAllSyntax,
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, pkg := range load {
		info := Info(pkg)
		_ = info
	}
}
