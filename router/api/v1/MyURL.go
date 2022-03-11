package v1

import (
	"URL_Shortener/logic"
	"URL_Shortener/model"
	"URL_Shortener/model/mysql"
	"URL_Shortener/model/redis"
	my_hashids "URL_Shortener/package/hashids"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type post_url struct {
	ExpireAt time.Time `json:"expireAt"`
	URL      string    `json:"url"`
}

func AddURL(c *gin.Context) {
	//m := make(map[string]interface{})

	var data_url post_url
	if err := c.BindJSON(&data_url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "參數解析錯誤"})
		return
	}

	if data_url.URL == "" || !logic.VerifyURL(data_url.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "url錯誤"})
		return
	}

	if data_url.ExpireAt.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "expireAt已過期"})
		return
	}

	//資料庫讀寫
	var hash_id string
	my_url, mysql_read_err := mysql.ReadByURL(data_url.URL)
	if mysql_read_err != gorm.ErrRecordNotFound && mysql_read_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
		return
	} else if mysql_read_err == gorm.ErrRecordNotFound {
		my_url := model.MyURL{URL: data_url.URL, ExpireAt: data_url.ExpireAt}
		id, err := mysql.Create(my_url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
			return
		}
		hash_id, err = my_hashids.NewHashID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
			return
		}
		err = redis.SetURL(hash_id, data_url.URL, data_url.ExpireAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
			return
		}
	} else {
		var err error
		hash_id, err = my_hashids.NewHashID(my_url.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
			return
		}
		if my_url.ExpireAt.Before(data_url.ExpireAt) {
			my_url.ExpireAt = data_url.ExpireAt
			err := mysql.Update(*my_url)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
				return
			}
			err = redis.UpdateURL(hash_id, data_url.ExpireAt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
				return
			}
		}
	}

	//回傳結果
	short_url := fmt.Sprintf("http://localhost:8080/%s", hash_id)
	c.JSON(http.StatusOK, gin.H{"id": hash_id, "short_url": short_url})

}

func GetURL(c *gin.Context) {
	str_id := c.Param("id")
	url, err := redis.GetURL(str_id)

	if err != nil {
		if err.Error() == "redigo: nil returned" {
			c.JSON(http.StatusNotFound, gin.H{"message": "連結不存在或已過期"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
