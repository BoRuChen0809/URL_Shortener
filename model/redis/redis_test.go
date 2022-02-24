package redisurl_test

import (
	"URL_Shortener/model"
	redisurl "URL_Shortener/model/redis"
	"testing"
	"time"
)

func Test_redis(t *testing.T) {
	test_url := model.MyURL{ExpireAt: time.Now().Add(time.Hour), URL: "test_url"}
	err := redisurl.Insert(test_url.ExpireAt.Format(time.RFC3339), "test", test_url.URL)
	if err != nil {
		t.Error(err)
	}

	url, err := redisurl.SelectByID("test")
	if err != nil {
		t.Error(err)
	}
	if url != test_url.URL {
		t.Errorf("result is not same, expect:%s, output:%s", test_url.URL, url)
	}

}
