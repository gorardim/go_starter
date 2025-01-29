package gormx

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CdnOrderAttr struct {
	Address string `json:"address"`
}

func (CdnOrderAttr) FieldName() string {
	return "attr"
}

type CdbOrder struct {
	OrderId string                   `json:"order_id"`
	Attr    JsonObject[CdnOrderAttr] `json:"attr"`
}

func TestJsonType_1(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		o := &CdbOrder{
			OrderId: "123",
			Attr: JsonObject[CdnOrderAttr]{
				Data: CdnOrderAttr{},
			},
		}
		marshal, err := json.Marshal(o)
		assert.NoError(t, err)
		t.Logf("marshal: %s", marshal)
	})

	t.Run("test", func(t *testing.T) {
		o := &CdbOrder{
			OrderId: "123",
		}

		marshal, err := json.Marshal(o)
		assert.NoError(t, err)
		t.Logf("marshal: %s", marshal)
	})

	t.Run("test", func(t *testing.T) {
		o := &CdbOrder{}
		err := json.Unmarshal([]byte(`{"attr":{"address": "123"}}`), o)
		assert.NoError(t, err)
	})

	t.Run("test", func(t *testing.T) {
		o := &CdbOrder{}
		err := json.Unmarshal([]byte(`{}`), o)
		assert.NoError(t, err)
	})
}
