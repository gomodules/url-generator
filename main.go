package main

import (
	"fmt"
	"net/url"

	flag "github.com/spf13/pflag"
)

func main() {
	var (
		scheme, host, user, passwd string
	)
	flag.StringVarP(&scheme, "scheme", "s", "", "URL scheme")
	flag.StringVarP(&host, "host", "h", "", "URL host")
	flag.StringVarP(&passwd, "user", "u", "", "URL username")
	flag.StringVarP(&passwd, "password", "p", "", "URL password")
	flag.Parse()

	var u url.URL
	u.Scheme = scheme
	u.Host = host
	u.User = url.UserPassword(user, passwd)
	fmt.Println(u.String())
}
