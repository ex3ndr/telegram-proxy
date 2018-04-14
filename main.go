package main

import (
	"github.com/armon/go-socks5"
)
import "os"

func main() {
	creds := socks5.StaticCredentials{os.Getenv("SOCKS_USER"): os.Getenv("SOCKS_PASSWORD")}
	cator := socks5.UserPassAuthenticator{Credentials: creds}
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{cator},
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe("tcp", "0.0.0.0:1080"); err != nil {
		panic(err)
	}
}
