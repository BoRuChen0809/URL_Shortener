package redis

import (
	"URL_Shortener/global"
	"time"

	"github.com/gomodule/redigo/redis"
)

func SetURL(id, url string, expireAt time.Time) error {
	c := global.Redis_Pool.Get()
	defer c.Close()

	c.Send("set", id, url)
	c.Send("expireat", id, expireAt.Unix())
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

func UpdateURL(id string, expireAt time.Time) error {
	c := global.Redis_Pool.Get()
	defer c.Close()

	_, err := c.Do("expireat", id, expireAt.Unix())
	return err
}
