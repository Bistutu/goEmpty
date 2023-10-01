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
	httputil.Get("https://jwxt.bistu.edu.cn/jwapp/i18n.do?appName=emaphome&EMAP_LANG=zh",
		http.Header{"Referer": []string{"https://jwxt.bistu.edu.cn/jwapp/sys/emaphome/portal/index.do"}})

	httputil.Get("https://jwxt.bistu.edu.cn/jwapp/sys/emaphome/appShow.do?id=e84984ccc396400085b3989e7819c063",
		http.Header{"Referer": []string{"https://jwxt.bistu.edu.cn/jwapp/sys/emaphome/portal/index.do"}})

	httputil.Get("https://jwxt.bistu.edu.cn/jwapp/sys/jwpubapp/commonUseApp/updateVisitAppCount.do?appName=kxjas&appId=4768402106681759",
		http.Header{"Referer": []string{"https://jwxt.bistu.edu.cn/jwapp/sys/kxjas/*default/index.do?EMAP_LANG=zh"}})

	httputil.Get("https://jwxt.bistu.edu.cn/jwapp/sys/funauthapp/api/getAppConfig/kxjas-4768402106681759.do?v=04011173702732529",
		http.Header{"Referer": []string{"https://jwxt.bistu.edu.cn/jwapp/sys/kxjas/*default/index.do?EMAP_LANG=zh"}})

	httputil.Get("https://jwxt.bistu.edu.cn/jwapp/sys/jwpubapp/pub/setJwCommonAppRole.do",
		http.Header{"Referer": []string{"https://jwxt.bistu.edu.cn/jwapp/sys/kxjas/*default/index.do?EMAP_LANG=zh"}})
	// 返回

	cookies = httputil.GetCookies("https://jwxt.bistu.edu.cn/jwapp/")
	// 将 cookie 暂时保存至文件
	file, _ := os.Create("cookie.txt")
	defer file.Close()
	file.WriteString(httputil.CookiesToString(cookies))
	return cookies, nil
}
