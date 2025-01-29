package jwtutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type user struct {
	Name string `json:"name"`
}

type user2 struct {
	Name2 string `json:"name2"`
}

func TestNew(t *testing.T) {
	data := &user{
		Name: "test",
	}
	token, err := Sign("", time.Hour, data)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	fmt.Println(token)
}

func TestParse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		data := &user{
			Name: "test",
		}
		token, err := Sign("123", time.Second, data)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		fmt.Println(token)

		claim, err := Parse[user]("123", token)
		assert.NoError(t, err)
		assert.Equal(t, data.Name, claim.Custom.Name)

	})

	t.Run("expired", func(t *testing.T) {
		data := &user2{
			Name2: "test2",
		}
		token, err := Sign("123", 0, data)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claim, err := Parse[user2]("123", token)
		assert.Error(t, err)
		assert.Nil(t, claim)
	})
}
