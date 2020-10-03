package gosms

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var config Config

// 初始化
func Init(input Config) {
	config = input
}

// 發送文字訊息
func SendText(mobile, message string) (bool, error) {

	// 獲取完整路徑
	fullURL := formatFullURL(mobile, message)

	resp, err := http.Get(fullURL)

	if err != nil {
		return false, nil
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, nil
	}

	log.Print(string(bytes))

	return true, nil
}

func formatFullURL(mobile, message string) string {
	baseURL := "https://api.twsms.com/json/sms_send.php"
	fullURL := fmt.Sprintf("%s?username=%s&password=%s", baseURL, config.Account, config.Password)
	fullURL  = fmt.Sprintf("%s&mobile=%s&message=%s", fullURL, mobile, message)

	return fullURL
}

type Config struct {
	Account  string
	Password string
}
