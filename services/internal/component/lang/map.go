package lang

import (
	"context"
	_ "embed"

	"app/api/model"

	"gopkg.in/yaml.v3"
)

//go:embed lang.yml
var langYml []byte

var langMap = map[string]model.Lang{}

func Init() error {
	if err := yaml.Unmarshal(langYml, &langMap); err != nil {
		return err
	}
	return nil
}

func T(ctx context.Context, key string) string {
	l := FromContext(ctx)
	v, ok := langMap[key]
	if !ok {
		return key
	}
	switch l {
	case "zh":
		return v.Zh
	case "mn":
		return v.Mn
	case "vn":
		return v.Vn
	case "es":
		return v.Es
	case "fr":
		return v.Fr
	case "zh_tr":
		return v.ZhTr
	case "ja":
		return v.Ja
	case "ko":
		return v.Ko
	case "pt":
		return v.Pt
	}
	return v.En
}
