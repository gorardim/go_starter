package model

// AdminImage
type AdminImage struct {
	Id   int    `json:"id" gorm:"primaryKey;column:id"` // Id
	Name string `json:"name" gorm:"column:name"`        // Name
	Url  string `json:"url" gorm:"column:url"`          // Url
}

func (*AdminImage) TableName() string {
	return "admin_image"
}

func (*AdminImage) PK() string {
	return "id"
}
