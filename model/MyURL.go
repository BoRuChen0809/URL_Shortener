package model

import "time"

type MyURL struct {
	ID       uint64 `gorm:"primary_key;"`
	ExpireAt time.Time
	URL      string
}

func TableName() string {
	return "my_urls"
}
