package redisurl_test

import (
	"URL_Shortener/model"
	redisurl "URL_Shortener/model/redis"
	"testing"
	"time"
)

func Test_redis(t *testing.T) {
	test_url := model.MyURL{ExpireAt: time.Now().Add(time.Hour), URL: "test_url"}
	err := redisurl.SetURL("test", test_url.URL, test_url.ExpireAt.Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	url, err := redisurl.GetURL("test")
	if err != nil {
		t.Error(err)
	}
	if url != test_url.URL {
		t.Errorf("result is not same, expect:%s, output:%s", test_url.URL, url)
	}
}

func Test_seturl(t *testing.T) {
	time := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	err := redisurl.SetURL("Test_id", "Test_URL", time)
	if err != nil {
		t.Error(err)
	}
}

func Test_geturl(t *testing.T) {
	url, err := redisurl.GetURL("Test_id")
	if err != nil {
		t.Error(err)
	}
	t.Log(url)
}

func Test_updateurl(t *testing.T) {
	time := time.Now().Add(100 * time.Minute).Format(time.RFC3339)
	err := redisurl.UpdateURL("Test_id", time)
	if err != nil {
		t.Error(err)
	}
}
