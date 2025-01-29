package api

import (
	"fmt"
	"testing"
)

func Test_genCustomId(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(genCustomId())
	}
}
