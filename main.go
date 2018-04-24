package main

import (
	"github.com/armon/go-socks5"
	"golang.org/x/net/context"
)
import (
	"os"
	"strings"
	"net"
)

type IpRules struct {
	ipAddresses []net.IP
	ipNetworks  []*net.IPNet
};

func (p *IpRules) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	for _, ip := range p.ipAddresses {
		if ip.Equal(req.DestAddr.IP) {
			return ctx, true
		}
	}

	for _, ipNet := range p.ipNetworks {
		if ipNet.Contains(req.DestAddr.IP) {
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

	if os.Getenv("SOCKS_WHITELIST") != "" {
		ipAddresses := strings.Split(os.Getenv("SOCKS_WHITELIST"), ",")
		rules := &IpRules{}
		for _, s := range ipAddresses {
			ipAddr, ipNet, err := net.ParseCIDR(s)
			if err != nil {
				ipAddr = net.ParseIP(s)
				if ipAddr == nil {
					panic("Wrong IP or Network address")
				}
				rules.ipAddresses = append(rules.ipAddresses, ipAddr)
				continue
			}

			rules.ipNetworks = append(rules.ipNetworks, ipNet)
		}
		conf.Rules = rules
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe("tcp", "0.0.0.0:1080"); err != nil {
		panic(err)
	}
}
