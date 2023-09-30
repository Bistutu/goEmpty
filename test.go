package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gocolly/colly"

	"GoEmpty/utils"
	"GoEmpty/utils/httputil"
)

func login(username, password string) (map[string]string, error) {
	var encryKey string

	c := colly.NewCollector()
	c.AllowURLRevisit = true
	// 忽略证书错误
	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})

	c.OnHTML("#pwdEncryptSalt", func(e *colly.HTMLElement) {
		encryKey = e.Attr("value") // 获取盐值
	})
	// 请求登录页
	err := c.Visit("https://wxjw.bistu.edu.cn/authserver/login?service=https://jwxt.bistu.edu.cn:443/jwapp/sys/emaphome/portal/index.do")
	if err != nil {
		return nil, err
	}

	password, err = utils.Encrypt(password, encryKey)
	if err != nil {
		return nil, err
	}
	// 至此，获得了本次账号和密码
	httputil.AddCookie("https://wxjw.bistu.edu.cn/authserver/login", c.Cookies("https://wxjw.bistu.edu.cn/authserver/login"))
	// 判断是否需要验证码：

	return nil, nil // return appropriate value
}

func main() {
	_, err := login("123456", "123456")
	if err != nil {
		log.Fatal(err)
	}
}
