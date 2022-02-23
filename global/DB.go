package global

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
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
	SetMySQL()
}

func SetMySQL() {
	var err error
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DBEngine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqldb, _ := DBEngine.DB()
	sqldb.SetConnMaxIdleTime(10)
	sqldb.SetMaxOpenConns(100)

}
