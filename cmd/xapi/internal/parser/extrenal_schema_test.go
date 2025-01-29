package parser

import (
	"encoding/json"
	"testing"
)

func Test_parseExternalSchema(t *testing.T) {
	schema := parseExternalSchema("/Users/huqi/waibao/travel/travel-backend-go/api/model")
	marshal, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	t.Log(string(marshal))
}
