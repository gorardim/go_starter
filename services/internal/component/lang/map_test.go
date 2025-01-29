package lang

import (
	"fmt"
	"slices"
	"testing"

	"app/api/model"
	"github.com/stretchr/testify/assert"
)

type langExt struct {
	model.Lang
	key string
}

func TestMap(t *testing.T) {
	var langArr []langExt
	for key, lang := range langMap {
		if lang.En != "" {
			langArr = append(langArr, langExt{lang, key})
		}
	}
	// sort
	slices.SortFunc(langArr, func(i, j langExt) int {
		if i.key < j.key {
			return -1
		}
		if i.key > j.key {
			return 1
		}
		return 0
	})

	for _, lang := range langArr {
		fmt.Printf("%s: \n", lang.key)
		fmt.Printf("  en: %s\n", lang.En)
		fmt.Printf("  zh: %s\n", lang.Zh)
	}
}

func TestInit(t *testing.T) {
	err := Init()
	assert.NoError(t, err)
	t.Log(langMap)
}
