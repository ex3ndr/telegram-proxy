package main

import (
	"github.com/armon/go-socks5"
)
import "os"

func main() {
	user := "user"
	password := "password"
	if os.Getenv("SOCKS_USER") != "" {
		user = os.Getenv("SOCKS_USER")
	}
	if os.Getenv("SOCKS_PASSWORD") != "" {
		password = os.Getenv("SOCKS_PASSWORD")
	}
	creds := socks5.StaticCredentials{user: password}
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
