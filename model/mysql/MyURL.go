package mysql

import (
	"URL_Shortener/global"
	"URL_Shortener/model"
)

func Create(myurl model.MyURL) (int64, error) {
	err := global.DBEngine.Table(myurl.TableName()).Create(&myurl).Error
	return myurl.ID, err
}

func ReadByID(id int64) (*model.MyURL, error) {
	myurl := model.MyURL{ID: id}
	err := global.DBEngine.Table(myurl.TableName()).Where("id = ?", id).
		First(&myurl).Error
	return &myurl, err
}

func ReadByURL(url string) (*model.MyURL, error) {
	myurl := model.MyURL{URL: url}
	err := global.DBEngine.Table(myurl.TableName()).Where("url = ?", url).
		First(&myurl).Error
	return &myurl, err
}

func Update(myurl model.MyURL) error {
	url := model.MyURL{ID: myurl.ID}
	return global.DBEngine.Table(url.TableName()).Where("id = ?", url.ID).
		First(&url).Updates(myurl).Error
}
