package randx

import (
	"fmt"
	"testing"
)

func TestSeq(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(Seq(10))
	}
}

func TestDigit(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(Digit(10))
	}
}

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(Int(100))
	}
}

func TestAlpha(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(Alpha(2))
	}
}
