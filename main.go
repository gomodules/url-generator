package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"net/url"
)

func main() {
	var passwd string
	flag.StringVar(&passwd, "token", "", "Zepto mail token")

	var u url.URL
	u.Scheme = "starttls"
	u.Host = "smtp.zeptomail.com:587"
	u.User = url.UserPassword("emailapikey", passwd)
	fmt.Println(u.String())
}
