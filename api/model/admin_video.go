package model

// AdminVideo
type AdminVideo struct {
	Id   int    `json:"id" gorm:"primaryKey;column:id"` // Id
	Name string `json:"name" gorm:"column:name"`        // Name
	Url  string `json:"url" gorm:"column:url"`          // Url
}

func (*AdminVideo) TableName() string {
	return "admin_video"
}

func (*AdminVideo) PK() string {
	return "id"
}
