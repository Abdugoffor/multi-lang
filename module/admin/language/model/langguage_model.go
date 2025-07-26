package language_model

import "time"

type Language struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug" gorm:"uniqueIndex"`
	IsActive  bool   `json:"is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (lang *Language) TableName() string {
	return "languages"
}
