package httputil

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
)

var client = &http.Client{
	Transport: &http.Transport{
		// 因为学校采用的是自签证书，所以这里需要跳过 SSL/TLS 证书验证
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}
var once sync.Once

func init() {
	once.Do(func() {
		jar, _ := cookiejar.New(nil)
		client.Jar = jar
	})
}
func Get(link string, header http.Header) (*http.Response, error) {
	req, _ := http.NewRequest("GET", link, nil)
	if header != nil {
		req.Header = header
	}
	req.Header.Add("User-Agent", UserAgent)
	// 暂时性的将请求打印出来
	response, _ := client.Do(req)
	all, _ := io.ReadAll(response.Body)
	fmt.Println(link + "===》》》" + string(all))
	return response, nil
}

func Post(link string, header http.Header, data io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest("POST", link, data)
	if header != nil {
		req.Header = header
	}
	req.Header.Add("User-Agent", UserAgent)
	return client.Do(req)
}

func PostForm(link string, header http.Header, params url.Values) (*http.Response, error) {
	req, _ := http.NewRequest("POST", link, strings.NewReader(params.Encode()))
	if header != nil {
		req.Header = header
	}
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(req)
}

func AddCookie(link string, cookies []*http.Cookie) {
	parse, _ := url.Parse(link)
	client.Jar.SetCookies(parse, cookies)
}
func GetCookies(link string) []*http.Cookie {
	parse, _ := url.Parse(link)
	return client.Jar.Cookies(parse)
}

func removeAllCookie() {
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
}

// CookiesToString 将 []*http.Cookie 转换为字符串。
func CookiesToString(cookies []*http.Cookie) string {
	cookieBytes, _ := json.Marshal(cookies)
	return string(cookieBytes)
}

func StringToCookie(data string) []*http.Cookie {
	cookies := make([]*http.Cookie, 8)
	err := json.Unmarshal([]byte(data), &cookies)
	if err != nil {
		log.Fatalf("read cookie file err: %v", err)
		return nil
	}
	return cookies
}
