package mysql_test

import (
	"URL_Shortener/model"
	"URL_Shortener/model/mysql"
	"testing"
	"time"
)

func Test_mysql(t *testing.T) {
	test_url := model.MyURL{ExpireAt: time.Now().Add(time.Hour), URL: "test_url"}
	id, err := mysql.Create(test_url)
	if err != nil {
		t.Error(err)
	}
	t.Log(id)

	myurl, err := mysql.ReadByID(id)
	if err != nil {
		t.Error(err)
	}
	if myurl.URL != test_url.URL {
		t.Errorf("result is not same, expect:%s, output:%s\n", test_url.URL, myurl.URL)
	}
	t.Log(myurl)

	myurl, err = mysql.ReadByURL(test_url.URL)
	if err != nil {
		t.Error(err)
	}
	if myurl.URL != test_url.URL {
		t.Errorf("result is not same, expect:%s, output:%s\n", test_url.URL, myurl.URL)
	}
	t.Log(myurl)

	myurl.URL = "Update_URL"
	myurl.ExpireAt = time.Now().Add(time.Hour * 10)
	err = mysql.Update(*myurl)
	if err != nil {
		t.Error(err)
	}
}

func Test_readbyID(t *testing.T) {
	myurl, err := mysql.ReadByID(0)
	if err != nil {
		t.Error(err)
	}
	t.Log(myurl)
}

func Test_readbyURL(t *testing.T) {
	myurl, err := mysql.ReadByURL("")
	if err != nil {
		t.Error(err)
	}
	t.Log(myurl)
}

func Test_update(t *testing.T) {
	myurl := model.MyURL{URL: "", ExpireAt: time.Now()}
	err := mysql.Update(myurl)
	if err != nil {
		t.Error(err)
	}
	t.Log(myurl)
}
