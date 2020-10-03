package gosms

import "fmt"

var config2 Config

func OOOO() {
	fmt.Print(9999)
}

func Initialization(config Config) {
	config2 = config


}

func SendText(mobile, message string) {

	aaa := fmt.Sprintf("https://api.twsms.com/json/sms_send.php?username=%s&password=%s")
	ssw := fmt.Sprintf("%s&mobile=%s&message=%s", aaa, mobile, message)

	fmt.Sprintf(ssw)
}

type Config struct {
	Account  string
	Password string
}
