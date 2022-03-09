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

func AddURL(c *gin.Context) {
	m := make(map[string]interface{})

	//JSON轉map
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "參數解析錯誤"})
		return
	}

	//處理url
	url := fmt.Sprint(m["url"])
	if url == "" || !logic.VerifyURL(url) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "url錯誤"})
		return
	}

	//處理expire_at
	str_time := fmt.Sprint(m["expireAt"])
	expire_at, err := time.Parse(time.RFC3339, str_time)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "expireAt解析錯誤"})
		return
	}
	if expire_at.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "expireAt已過期"})
		return
	}

	//資料庫讀寫
	var hash_id string
	my_url, mysql_read_err := mysql.ReadByURL(url)
	if mysql_read_err != gorm.ErrRecordNotFound && mysql_read_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
		return
	} else if mysql_read_err == gorm.ErrRecordNotFound {
		my_url := model.MyURL{URL: url, ExpireAt: expire_at}
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
		err = redis.SetURL(hash_id, url, str_time)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
			return
		}
	} else {
		hash_id, err = my_hashids.NewHashID(my_url.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
			return
		}
		if my_url.ExpireAt.Before(expire_at) {
			my_url.ExpireAt = expire_at
			err := mysql.Update(*my_url)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
				return
			}
			err = redis.UpdateURL(hash_id, str_time)
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
