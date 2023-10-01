package main

import (
	"fmt"
	"io"
	"os"

	"GoEmpty/utils/httputil"
)

func main() {
	//cookies, err := loginutil.Login("2018011264", "aa286838")
	//cookies, err := loginutil.GetEmptyCookie("2018011264", "aa286838")
	//cookies, err := loginutil.GetEmptyCookie("2018011170", "aa100018")
	//cookies, err := loginutil.GetEmptyCookie("2018011184", "aaa212265")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("final cookie: %v", cookies)
	open, _ := os.Open("cookie.txt")
	all, _ := io.ReadAll(open)
	cookie := httputil.StringToCookie(string(all))
	for k, v := range cookie {
		fmt.Println(k, v)
	}
	fmt.Println(cookie)
}
