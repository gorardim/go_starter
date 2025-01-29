package ginx

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Desensitize map[request path] []desensitize json path
// eg: map["/api/v1/user/login"] = []string{"password"}
type Desensitize map[string][]string

func (d Desensitize) Desensitize(requestPath string, body []byte) []byte {
	if len(d) == 0 {
		return body
	}
	if path, ok := d[requestPath]; ok {
		return desensitize(path, body)
	}
	return body
}

func desensitize(path []string, body []byte) []byte {
	g := gjson.ParseBytes(body)
	for _, p := range path {
		if !g.Get(p).Exists() {
			continue
		}
		body, _ = sjson.SetBytes(body, p, "******")
	}
	return body
}
