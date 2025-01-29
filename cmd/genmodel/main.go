package main

import (
	_ "embed"

	"app/cmd/genmodel/internal"
)

func main() {
	internal.Generate(&internal.Options{
		TypeMapping: map[string]string{
			"article.title":               "LangType",
			"article.content":             "LangType",
			"article_category.name":       "LangType",
			"invest_contract.title":       "LangType",
			"invest_order.contract_title": "LangType",
			"ad_position.link":            "LinkType",
			"ad_position.title":           "LangType",
			"ad_king_kong.name":           "LangType",
			"ad_king_kong.link":           "LinkType",
			"travel.name":                 "LangType",
			"travel.content":              "LangType",
			"task.config":                 "TaskConfigAttrType",
			"task.title":                  "LangType",
			"task.sub_title":              "LangType",
			"task.content":                "LangType",
		},
		EnumTypes: []string{
			"ad_position.code",
		},
	})
}
