package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"net/url"
)

func main() {
	var passwd string
	flag.StringVarP(&passwd, "password", "p", "", "Zepto mail token")
	flag.Parse()

	var u url.URL
	u.Scheme = "starttls"
	u.Host = "smtp.zeptomail.com:587"
	u.User = url.UserPassword("emailapikey", passwd)
	fmt.Println(u.String())
}
