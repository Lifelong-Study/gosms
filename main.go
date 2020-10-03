package gosms

import (
	"fmt"
	"log"
)

var config Config

func Initialization(input Config) {
	config = input

}

func SendText(mobile, message string) {

	aaa := fmt.Sprintf("https://api.twsms.com/json/sms_send.php?username=%s&password=%s", config.Account, config.Password)
	ssw := fmt.Sprintf("%s&mobile=%s&message=%s", aaa, mobile, message)

	log.Print(ssw)
}

type Config struct {
	Account  string
	Password string
}
