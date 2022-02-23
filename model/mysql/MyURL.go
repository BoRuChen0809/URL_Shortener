package mysql

import (
	"URL_Shortener/global"
	"URL_Shortener/model"

	"time"
)

func Insert(expireAt time.Time, url string) (uint64, error) {
	myurl := model.MyURL{ExpireAt: expireAt, URL: url}
	err := global.DBEngine.Table(model.TableName()).Create(&myurl).Error
	return myurl.ID, err
}

func SelectByID(id uint64) (*model.MyURL, error) {
	myurl := model.MyURL{ID: id}
	err := global.DBEngine.Table(model.TableName()).Where(
		"id = ? AND expire_at > ?", id, time.Now()).First(&myurl).Error
	return &myurl, err
}
