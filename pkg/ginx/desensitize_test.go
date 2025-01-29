package ginx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_desensitize(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		bytes := desensitize([]string{
			"password",
		}, []byte(`{"password":"123456"}`))
		assert.Equal(t, `{"password":"******"}`, string(bytes))
	})

	t.Run("bad json", func(t *testing.T) {
		bytes := desensitize([]string{
			"password",
		}, []byte(`123456`))
		assert.Equal(t, `123456`, string(bytes))
	})

}
