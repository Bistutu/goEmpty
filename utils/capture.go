package utils

import (
	"bytes"
	"image"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/anthonynsimon/bild/effect"
	"github.com/otiai10/gosseract/v2"

	"GoEmpty/utils/httputil"
)

func getCaptureCode() string {
	url := "https://wxjw.bistu.edu.cn/authserver/getCaptcha.htl?" + strconv.Itoa(int(time.Now().UnixMilli()))
	resp, err := httputil.GET(url, nil)
	if err != nil {
		log.Fatalf("无法下载图片: %v", err)
	}
	defer resp.Body.Close()

	imgData, err := io.ReadAll(resp.Body)

	file2, _ := os.Create("capture2.png")
	defer file2.Close()
	file2.Read(imgData)

	// 图片去噪
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		log.Fatalf("image decode err: %v", err)
		return ""
	}
	gray := effect.Grayscale(img)

	// 将图片写至本地，看看准确度
	file, _ := os.Create("capture.png")
	defer file.Close()
	file.Read(gray.Pix)

	if err != nil {
		log.Fatalf("无法读取图片数据: %v", err)
	}

	client := gosseract.NewClient()
	defer client.Close()
	client.SetImageFromBytes(imgData)

	code, err := client.Text()
	if err != nil {
		log.Fatalf("无法识别验证码: %v", err)
	}
	return strings.TrimSpace(code)
}

// 判断是否需要验证码
func isNeedCaptcha(username string) bool {
	resp, err := httputil.GET("https://wxjw.bistu.edu.cn/authserver/checkNeedCaptcha.htl?username="+username, nil)
	if err != nil {
		log.Fatalf("get isNeedCapture err: %v", err)
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("readAll Body err: %v", err)
		return false
	}
	if strings.Contains(string(body), "true") {
		return true
	}
	return false
}
