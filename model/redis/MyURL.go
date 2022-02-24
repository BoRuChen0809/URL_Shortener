package redisurl

import (
	"URL_Shortener/global"
	"time"

	"github.com/gomodule/redigo/redis"
)

func Insert(expireAt, id, url string) error {
	c := global.Redis_Pool.Get()
	defer c.Close()

	if expireAt == "" {
		c.Send("set", id, "")
		c.Send("expire", id, 1000)
	} else {
		c.Send("set", id, url)
		t, _ := time.Parse(time.RFC3339, expireAt)
		if t.After(time.Now().Add(time.Hour)) {
			c.Send("expire", id, 3600)
		} else {
			c.Send("expireat", id, t.Unix())
		}
	}

	c.Flush()

	_, err := c.Receive()
	if err != nil {
		return err
	}
	return nil
}

func SelectByID(id string) (string, error) {
	c := global.Redis_Pool.Get()
	defer c.Close()
	url, err := redis.String(c.Do("get", id))
	if err != nil {
		return "", err
	}
	return url, err
}
