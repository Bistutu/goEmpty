package loginutil

import (
	"net/http"
	"os"

	"GoEmpty/utils/httputil"
)

func GetEmptyCookie(username, password string) ([]*http.Cookie, error) {
	cookies, err := Login(username, password)
	if err != nil {
		return nil, err
	}
	// 开始发送 5 次请求
	// 1
	httputil.Post("https://jwxt.bistu.edu.cn/jwapp/sys/jwpubapp/gjhwhController/queryAppGjhMapping.do",
		http.Header{"Referer": []string{"https://jwxt.bistu.edu.cn/jwapp/sys/kxjas/*default/index.do?EMAP_LANG=zh"}}, nil)
	// 2
	httputil.Post("https://jwxt.bistu.edu.cn/jwapp/sys/jwpubapp/modules/bb/cxjwggbbdqx.do",
		http.Header{"Referer": []string{"https://jwxt.bistu.edu.cn/jwapp/sys/kxjas/*default/index.do?EMAP_LANG=zh"}}, nil)
	// 3
	httputil.Get("https://jwxt.bistu.edu.cn/jwapp/sys/funauthapp/api/getAppConfig/kxjas-4768402106681759.do", nil)
	// 4
	httputil.Get("https://jwxt.bistu.edu.cn/jwapp/sys/emappagelog/config/kxjas.do", nil)
	// 5
	httputil.Post("https://jwxt.bistu.edu.cn/jwapp/sys/emapcomponent/schema/getList.do",
		http.Header{"Referer": []string{"https://jwxt.bistu.edu.cn/jwapp/sys/kxjas/*default/index.do"}}, nil)
	// 返回
	cookies = httputil.GetCookies("https://jwxt.bistu.edu.cn/")
	// TODO 测试：将 cookie 暂时保存至文件
	file, _ := os.Create("cookie.txt")
	defer file.Close()
	file.WriteString(httputil.CookiesToString(cookies))
	return cookies, nil
}
