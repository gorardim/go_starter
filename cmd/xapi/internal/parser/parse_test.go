package parser

import (
	"fmt"
	"testing"
)

func TestParseDir(t *testing.T) {
	dir, err := ParseDir("/Users/huqi/waibao/repo-qsh-go/api/admin")

	fmt.Println(dir, err)
}
