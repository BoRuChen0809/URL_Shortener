package logic

import (
	"net/http"
	"net/url"
)

func VerifyURL(path string) bool {
	u, err := url.Parse(path)
	if err != nil {
		return false
	}
	switch u.Host {
	case "":
		return false
	case "localhost:8080":
		return false
	case "127.0.0.1:8080":
		return false
	}

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return false
	}
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	if res.StatusCode > 299 || res.StatusCode < 200 {
		return false
	}

	return true
}
