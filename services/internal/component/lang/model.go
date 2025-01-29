package lang

import (
	"context"

	"app/api/model"
)

func FromLangType(ctx context.Context, langType model.LangType) string {
	l := FromContext(ctx)
	switch l {
	case "en":
		return langType.Data.En
	case "zh":
		return langType.Data.Zh
	case "mn":
		if langType.Data.Mn != "" {
			return langType.Data.Mn
		}
	case "vn":
		if langType.Data.Vn != "" {
			return langType.Data.Vn
		}
	case "es":
		if langType.Data.Es != "" {
			return langType.Data.Es
		}
	case "fr":
		if langType.Data.Fr != "" {
			return langType.Data.Fr
		}
	case "zh_tr":
		if langType.Data.ZhTr == "" {
			return langType.Data.Zh
		}
		return langType.Data.ZhTr
	case "ja":
		if langType.Data.Ja != "" {
			return langType.Data.Ja
		}
	case "ko":
		if langType.Data.Ko != "" {
			return langType.Data.Ko
		}
	case "pt":
		if langType.Data.Pt != "" {
			return langType.Data.Pt
		}
	}
	return langType.Data.En
}

func FromLang(ctx context.Context, langType model.Lang) string {
	l := FromContext(ctx)
	switch l {
	case "en":
		return langType.En
	case "zh":
		return langType.Zh
	case "mn":
		if langType.Mn != "" {
			return langType.Mn
		}
	case "vn":
		if langType.Vn != "" {
			return langType.Vn
		}
	case "es":
		if langType.Es != "" {
			return langType.Es
		}
	case "fr":
		if langType.Fr != "" {
			return langType.Fr
		}
	case "zh_tr":
		if langType.ZhTr == "" {
			return langType.Zh
		}
		return langType.ZhTr
	case "ja":
		if langType.Ja != "" {
			return langType.Ja
		}
	case "ko":
		if langType.Ko != "" {
			return langType.Ko
		}
	case "pt":
		if langType.Pt != "" {
			return langType.Pt
		}
	}
	return langType.En
}

func ReturnLinkLang(ctx context.Context, link string) string {
	l := FromContext(ctx)
	switch l {
	case "en":
		return link + "?lang=" + l
	case "zh":
		return link + "?lang=" + l
	}
	return link + "?lang=en"
}
