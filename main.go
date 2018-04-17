package main

import (
	"github.com/armon/go-socks5"
	"golang.org/x/net/context"
)
import (
	"os"
	"strings"
)

type PermitCommand struct {
	ipAddresses []string
};

func (p *PermitCommand) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	for _, b := range p.ipAddresses {
		if b == req.DestAddr.IP.String() {
			return ctx, true
		}
	}

	return ctx, false
};

func main() {
	conf := &socks5.Config{}

	if os.Getenv("SOCKS_AUTH") != "no" {
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
		conf.AuthMethods = []socks5.Authenticator{cator}
	}

	if os.Getenv("IP_ADDRESSES") != "" {
		ipAddresses := strings.Split(os.Getenv("IP_ADDRESSES"), ",")
		conf.Rules = &PermitCommand{
			ipAddresses: ipAddresses,
		}
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe("tcp", "0.0.0.0:1080"); err != nil {
		panic(err)
	}
}
