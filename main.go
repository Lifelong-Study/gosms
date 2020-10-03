package gosms

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var config Config

type Platform uint8

const (
	PLATFORM_TWSMS Platform = iota
	PLATFORM_SMSGO
)

//
type Config struct {
	TWSMS Profile
	SMSGO Profile

	Platform Platform
}

type Profile struct {
	Account  string
	Password string
}

// 初始化
func Init(input Config) *Config {
	config = input

	return &config
}

//
func (config *Config) SetDefaultPlatform(platform Platform) *Config {
	config.Platform = platform
	return config
}

//
func (config *Config) UseTaiwanSMS() *Config {
	config.Platform = PLATFORM_TWSMS
	return config
}

//
func (config *Config) UseSmsGo() *Config {
	config.Platform = PLATFORM_SMSGO
	return config
}

// 發送文字訊息
func (config *Config) SendText(mobile, message string) (string, error) {

	if false {
	} else if config.Platform == PLATFORM_TWSMS {
		return config.useTaiwanSMSSendText(mobile, message)
	} else if config.Platform == PLATFORM_SMSGO {
		return config.useSmsGoSendText(mobile, message)
	}

	return "", errors.New("SMS platform set wrong")
}

//
func (config *Config) useTaiwanSMSSendText(mobile, message string) (string, error) {

	// 獲取完整路徑
	fullURL := formatTWSMSFullURL(mobile, message)

	resp, err := http.Get(fullURL)

	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	// {"code":"00000","text":"Success","msgid":338982447}

	return string(bytes), nil
}

//
func (config *Config) useSmsGoSendText(mobile, message string) (string, error) {
	return "", nil
}

// 私有函式: 組合完整路徑
func formatTWSMSFullURL(mobile, message string) string {
	baseURL := "https://api.twsms.com/json/sms_send.php"
	fullURL := fmt.Sprintf("%s?username=%s&password=%s", baseURL, config.TWSMS.Account, config.TWSMS.Password)
	fullURL  = fmt.Sprintf("%s&mobile=%s&message=%s", fullURL, mobile, message)

	return fullURL
}





func tt() {
	config := Config{
		TWSMS: Profile{ "",  ""},
		SMSGO: Profile{"",""},
	}

	sms := Init(config)

	sms.SetDefaultPlatform(PLATFORM_TWSMS)

	sms.
		SendText("0963265781", "test")
}





