package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFen2Yuan(t *testing.T) {
	assert.Equal(t, "1.67", Fen2Yuan(167))
	assert.Equal(t, "1.99", Fen2Yuan(199))
	assert.Equal(t, "0.00", Fen2Yuan(0))
}

func TestYuan2Fen(t *testing.T) {
	fen, err := Yuan2Fen("1.67")
	assert.Nil(t, err)
	assert.Equal(t, 167, fen)

	fen, err = Yuan2Fen("1.99")
	assert.Nil(t, err)
	assert.Equal(t, 199, fen)
}

func TestFormatAmount(t *testing.T) {
	assert.Equal(t, 1.3, FormatAmount(1.333))
	assert.Equal(t, 90, FormatAmount(90.00))
}
