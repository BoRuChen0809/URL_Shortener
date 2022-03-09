package redis

import (
	"URL_Shortener/global"
	"time"

	"github.com/gomodule/redigo/redis"
)

func SetURL(id, url, expireAt string) error {
	c := global.Redis_Pool.Get()
	defer c.Close()

	c.Send("set", id, url)
	t, _ := time.Parse(time.RFC3339, expireAt)
	c.Send("expireat", id, t.Unix())
	c.Flush()

	_, err := c.Receive()
	return err
}

func GetURL(id string) (string, error) {
	c := global.Redis_Pool.Get()
	defer c.Close()
	url, err := redis.String(c.Do("get", id))
	if err != nil {
		return "", err
	}
	return url, err
}

func UpdateURL(id, expireAt string) error {
	c := global.Redis_Pool.Get()
	defer c.Close()

	t, _ := time.Parse(time.RFC3339, expireAt)
	_, err := c.Do("expireat", id, t.Unix())
	return err
}
