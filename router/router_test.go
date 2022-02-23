package router_test

import (
	"URL_Shortener/model"
	"URL_Shortener/router"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

type Post_Case struct {
	Test_url    model.MyURL
	Expect_code int
}

var r *gin.Engine

func init() {
	r = router.NewRouter()
}

func Test_router(t *testing.T) {
	test_table := []Post_Case{
		{
			Expect_code: http.StatusBadRequest,
		},
		{
			Test_url: model.MyURL{
				URL:      "",
				ExpireAt: time.Now().Add(time.Hour),
			},
			Expect_code: http.StatusBadRequest,
		},
		{
			Test_url: model.MyURL{
				URL:      "https://www.google.com/",
				ExpireAt: time.Now(),
			},
			Expect_code: http.StatusBadRequest,
		},
		{
			Test_url: model.MyURL{
				URL:      "https://www.google.com/",
				ExpireAt: time.Now().Add(time.Hour),
			},
			Expect_code: http.StatusOK,
		},
	}

	for _, test_case := range test_table {
		w := httptest.NewRecorder()
		data, err := json.Marshal(test_case.Test_url)
		if err != nil {
			t.Error(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/api/mysql/urls", bytes.NewBuffer(data))
		if err != nil {
			t.Error(err)
		}
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		res_body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}
		if res.StatusCode != test_case.Expect_code {
			t.Errorf("status code is not same, expect:%d, output:%d\n", test_case.Expect_code, res.StatusCode)
		}

		t.Logf("%d : %s", res.StatusCode, res_body)
	}
}
