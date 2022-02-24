package router_test

import (
	"URL_Shortener/model"
	mypkg "URL_Shortener/package"
	"URL_Shortener/router"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
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

type Get_Case struct {
	ID          uint64
	Expect_code int
}

var (
	r *gin.Engine
	e *mypkg.Encodeing
)

func init() {
	r = router.NewRouter()
	e = mypkg.GetEncodeing()
	e.Set_encode()
	e.Set_decode()

}

func Test_PostV1(t *testing.T) {
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
func Test_GetV1(t *testing.T) {
	test_table := []Get_Case{
		{0, 404},
		{1, 301},
	}

	for _, test_case := range test_table {
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/mysql/%s", mypkg.ByteArray2String(e.Uint2ByteArray(test_case.ID)))
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Error(err)
		}

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

func Test_PostV2(t *testing.T) {
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

		req, err := http.NewRequest(http.MethodPost, "/api/v1/urls", bytes.NewBuffer(data))
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
func Test_GetV2(t *testing.T) {
	test_table := []Get_Case{
		{0, 404},
		{1, 301},
	}

	for _, test_case := range test_table {
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/%s", mypkg.ByteArray2String(e.Uint2ByteArray(test_case.ID)))
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Error(err)
		}

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

func Benchmark_PostV1(b *testing.B) {
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
				ExpireAt: time.Now().Add(time.Minute * -1),
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
		{
			Test_url: model.MyURL{
				URL:      "https://www.dcard.tw/f",
				ExpireAt: time.Now().Add(time.Hour),
			},
			Expect_code: http.StatusOK,
		},
	}

	b.SetParallelism(100000)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			case_id := rand.Intn(len(test_table))

			w := httptest.NewRecorder()
			data, err := json.Marshal(test_table[case_id].Test_url)
			if err != nil {
				b.Error(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/api/mysql/urls", bytes.NewBuffer(data))
			if err != nil {
				b.Error(err)
			}
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()
			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				b.Error(err)
			}
			if res.StatusCode != test_table[case_id].Expect_code {
				b.Errorf("status code is not same, expect:%d, output:%d\n", test_table[case_id].Expect_code, res.StatusCode)
			}
		}
	})
}
func Benchmark_GetV1(b *testing.B) {
	test_table := []Get_Case{
		{1000000000000000, 404},
		{1, 301},
		{2, 301},
		{0, 404},
		{3, 301},
		{4, 301},
		{5, 301},
		{6, 301},
		{7, 301},
		{8, 301},
	}

	b.SetParallelism(16)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			case_id := rand.Intn(len(test_table))
			w := httptest.NewRecorder()
			url := fmt.Sprintf("/mysql/%s", mypkg.ByteArray2String(e.Uint2ByteArray(test_table[case_id].ID)))
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				b.Error(err)
			}

			r.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()

			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				b.Error(err)
			}
			if res.StatusCode != test_table[case_id].Expect_code {
				b.Errorf("status code is not same, expect:%d, output:%d\n", test_table[case_id].Expect_code, res.StatusCode)
			}
		}
	})
}

func Benchmark_PostV2(b *testing.B) {
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
				ExpireAt: time.Now().Add(time.Hour * -1),
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
		{
			Test_url: model.MyURL{
				URL:      "https://www.dcard.tw/f",
				ExpireAt: time.Now().Add(time.Hour),
			},
			Expect_code: http.StatusOK,
		},
	}

	b.SetParallelism(100000)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			rand.Seed(time.Now().Unix())
			case_id := rand.Intn(len(test_table))

			w := httptest.NewRecorder()
			data, err := json.Marshal(test_table[case_id].Test_url)
			if err != nil {
				b.Error(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/api/v1/urls", bytes.NewBuffer(data))
			if err != nil {
				b.Error(err)
			}
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()
			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				b.Error(err)
			}
			if res.StatusCode != test_table[case_id].Expect_code {
				b.Errorf("status code is not same, expect:%d, output:%d\n", test_table[case_id].Expect_code, res.StatusCode)
			}
		}
	})
}
func Benchmark_GetV2(b *testing.B) {
	test_table := []Get_Case{
		{1000000000000000, 404},
		{1, 301},
		{2, 301},
		{0, 404},
		{3, 301},
		{4, 301},
		{5, 301},
		{6, 301},
		{7, 301},
		{8, 301},
	}

	b.SetParallelism(16)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			case_id := rand.Intn(len(test_table))
			w := httptest.NewRecorder()
			url := fmt.Sprintf("/%s", mypkg.ByteArray2String(e.Uint2ByteArray(test_table[case_id].ID)))
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				b.Error(err)
			}

			r.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()

			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				b.Error(err)
			}
			if res.StatusCode != test_table[case_id].Expect_code {
				b.Errorf("status code is not same, expect:%d, output:%d\n", test_table[case_id].Expect_code, res.StatusCode)
			}
		}
	})
}
