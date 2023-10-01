package main

import (
	"fmt"
	"log"

	"GoEmpty/utils/loginutil"
)

func main() {
	cookies, err := loginutil.Login("2018011264", "aa286838")
	//cookies, err := loginutil.GetEmptyCookie("2018011264", "aa286838")
	//cookies, err := loginutil.GetEmptyCookie("2018011170", "aa100018")
	//cookies, err := loginutil.GetEmptyCookie("2018011184", "aaa212265")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("final cookie: %v", cookies)
}
