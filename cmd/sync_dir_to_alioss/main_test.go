package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_walkDir(t *testing.T) {
	err := doSync(nil, "/Users/huqi/waibao/travel/travel-backend-go", "/")
	assert.NoError(t, err)
}
