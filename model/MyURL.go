package model

import "time"

type MyURL struct {
	ID       uint64    `gorm:"primary_key;"`
	ExpireAt time.Time `json:"expireAt"`
	URL      string    `json:"url"`
}

func TableName() string {
	return "my_urls"
}
