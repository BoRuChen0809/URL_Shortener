package model

import "time"

type MyURL struct {
	ID       int64    `gorm:"primary_key;"`
	ExpireAt time.Time `json:"expireAt"`
	URL      string    `json:"url"`
}

func (myurl MyURL) TableName() string {
	return "my_urls"
}
