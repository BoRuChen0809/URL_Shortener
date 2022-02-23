package mysql_test

import (
	"URL_Shortener/model/mysql"
	"testing"
	"time"
)

func Test_mysql(t *testing.T) {
	test_time := time.Now()
	url := "test_url"
	id, err := mysql.Insert(test_time, url)
	if err != nil {
		t.Error(err)
	}

	myurl, err := mysql.SelectByID(id)
	if err != nil {
		t.Error(err)
	}
	if myurl.URL != url {
		t.Errorf("result is not same, expect:%s, output:%s\n", url, myurl.URL)
	}
}
