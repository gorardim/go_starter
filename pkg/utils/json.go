package utils

import (
	"bytes"
	"encoding/json"
)

func JsonEncode(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func JsonEncodeByte(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func JsonEncodeBytePretty(v interface{}) []byte {
	b, _ := json.MarshalIndent(v, "", "  ")
	return b
}

func JsonEncodeUnEscape(v interface{}) string {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(v)
	return buf.String()
}
