package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordScanner_NextWord(t *testing.T) {
	t.Run("next word1", func(t *testing.T) {
		scanner := NewWordScanner("//x:topic charge.device.sync.status 设备状态同步")
		assert.Equal(t, "//x:topic", scanner.NextWord())
		assert.Equal(t, "charge.device.sync.status", scanner.NextWord())
		assert.Equal(t, "设备状态同步", scanner.NextWord())
	})

	t.Run("next word2", func(t *testing.T) {
		scanner := NewWordScanner("//x:channel default 设备状态同步 ")
		assert.Equal(t, "//x:channel", scanner.NextWord())
		assert.Equal(t, "default", scanner.NextWord())
		assert.Equal(t, "设备状态同步 ", scanner.Rest())
	})
}
