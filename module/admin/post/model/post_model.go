package post_model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type JSONBMap map[string]string

func (j JSONBMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONBMap) Scan(value interface{}) error {

	bytes, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	return json.Unmarshal(bytes, j)
}

type Post struct {
	ID          int64     `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Title       JSONBMap  `json:"title" gorm:"type:jsonb"`
	Description JSONBMap  `json:"description" gorm:"type:jsonb"`
	Content     JSONBMap  `json:"content" gorm:"type:jsonb"`
	View        int64     `json:"view" gorm:"default:0"`
	Slug        string    `json:"slug"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Post) TableName() string {
	return "posts"
}
