package global

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBEngine   *gorm.DB
	Redis_Pool *redis.Pool
)

const (
	USERNAME = "test"
	PASSWORD = "1234"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "shorturl"
)

func init() {
	SetupMySQL()
	SetupRedis()
}

func SetupMySQL() {
	var err error
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DBEngine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		panic(err)
	}

	sqldb, _ := DBEngine.DB()
	sqldb.SetConnMaxIdleTime(10)
	sqldb.SetMaxOpenConns(100)
}

func SetupRedis() {
	Redis_Pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 30 * time.Millisecond,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
