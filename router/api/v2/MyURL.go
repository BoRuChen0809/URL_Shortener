package v2

import (
	"URL_Shortener/global"
	"URL_Shortener/model/mysql"
	redisurl "URL_Shortener/model/redis"
	mypkg "URL_Shortener/package"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AddURL(c *gin.Context) {
	m := make(map[string]interface{})

	//JSON轉map
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "參數解析錯誤"})
		return
	}

	//處理url
	url := fmt.Sprint(m["url"])
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "url解析錯誤"})
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

	//寫入Mysql
	id, err := mysql.Insert(expire_at, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
		return
	}
	//編碼、寫入redis
	str_id := mypkg.ByteArray2String(global.Encode.Uint2ByteArray(id))
	err = redisurl.Insert(str_time, str_id, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "發生內部錯誤"})
		return
	}

	//回傳結果
	short_url := fmt.Sprintf("http://localhost:8080/%s", str_id)
	c.JSON(http.StatusOK, gin.H{"id": str_id, "short_url": short_url})
}

func GetURL(c *gin.Context) {
	str_id := c.Param("id")

	url, err := redisurl.SelectByID(str_id)
	if err != nil {
		//解碼
		id, err := global.Encode.ByteArray2Uint(mypkg.String2ByteArray(str_id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "解析url_id錯誤"})
			return
		}

		//查詢Mysql
		myurl, err := mysql.SelectByID(id)
		if err != nil {
			redisurl.Insert("", str_id, "")
			c.JSON(http.StatusNotFound, gin.H{"message": "連結不存在或已過期"})
			return
		} else {
			redisurl.Insert(myurl.ExpireAt.Format(time.RFC3339), str_id, myurl.URL)
			c.Redirect(http.StatusMovedPermanently, myurl.URL)
			return
		}
	} else if url == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "連結不存在或已過期"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
