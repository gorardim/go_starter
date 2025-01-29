package model

import "app/pkg/gormx"

const (
	Yes = "Y"
	No  = "N"
)

func YesOrNo(b bool) string {
	if b {
		return Yes
	}
	return No
}

const (
	MediaTypeVideo = "VIDEO"
	MediaTypeImage = "IMAGE"
)

type Lang struct {
	En   string `json:"en,omitempty"`
	Zh   string `json:"zh,omitempty"`
	Es   string `json:"es,omitempty"`
	Fr   string `json:"fr,omitempty"`
	Mn   string `json:"mn,omitempty"`
	Vn   string `json:"vn,omitempty"`
	Ja   string `json:"ja,omitempty"`
	Ko   string `json:"ko,omitempty"`
	Pt   string `json:"pt,omitempty"`
	ZhTr string `json:"zh_tr,omitempty"`
}

func (l Lang) FieldName() string {
	return ""
}

type LangType = gormx.JsonObject[Lang]

type LangObject struct {
	En   []*DisplaySettingItem `json:"en,omitempty"`
	Zh   []*DisplaySettingItem `json:"zh,omitempty"`
	Es   []*DisplaySettingItem `json:"es,omitempty"`
	Fr   []*DisplaySettingItem `json:"fr,omitempty"`
	Mn   []*DisplaySettingItem `json:"mn,omitempty"`
	Vn   []*DisplaySettingItem `json:"vn,omitempty"`
	Ja   []*DisplaySettingItem `json:"ja,omitempty"`
	Ko   []*DisplaySettingItem `json:"ko,omitempty"`
	Pt   []*DisplaySettingItem `json:"pt,omitempty"`
	ZhTr []*DisplaySettingItem `json:"zh_tr,omitempty"`
}

type DisplaySettingItem struct {
	// Title
	Title string `json:"title"`
	// Icon
	Icon string `json:"icon"`
	// desc
	Desc string `json:"desc"`
}

func (l LangObject) FieldName() string {
	return ""
}

type LangObjectType = gormx.JsonObject[LangObject]

type Link struct {
	// link type: ARTICLE_DETAIL_PAGE,ARTICLE_VIDEO_PAGE,DEPOSIT,INVEST,INVITE_FRIEND, HELP_CENTER,TASK,LEARN,MY_TEAM,SERVICE,COMING_SOON
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

func (l Link) FieldName() string {
	return ""
}

type LinkType = gormx.JsonObject[Link]

type Media struct {
	// media type: IMAGE,VIDEO
	Type string `json:"type,omitempty"`
	Src  string `json:"src,omitempty"`
}

type MediaType = gormx.JsonArray[Media]

type FixedTopDiscoveryHomePage struct {
	// ID
	Id int `json:"id"`
	// TYPE: POST, TRAVEL, NEWS, VIDEO
	Type string `json:"type"`
}
