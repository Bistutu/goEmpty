package httputil

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
)

var client = &http.Client{}

func init() {
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
}

func POST(link string, data io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest("POST", link, data)
	req.Header.Add("User-Agent", UserAgent)
	return client.Do(req)
}

func GET(link string, header http.Header) (*http.Response, error) {
	req, _ := http.NewRequest("GET", link, nil)
	if header != nil {
		req.Header = header
		req.Header.Add("User-Agent", UserAgent)
	}
	req.Header.Add("User-Agent", UserAgent)
	return client.Do(req)
}

func AddCookie(link string, cookies []*http.Cookie) {
	parse, _ := url.Parse(link)
	client.Jar.SetCookies(parse, cookies)
}
