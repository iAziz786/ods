package ips

import (
	"net/netip"
)

func CidrRange(cidr string) <-chan netip.Addr {
	prefix, err := netip.ParsePrefix(cidr)
	if err != nil {
		panic(err)
	}
	addrChan := make(chan netip.Addr)

	go func() {
		for addr := prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
			if addr.IsPrivate() {
				// skip the 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, and fc00::/7 address
				continue
			}
			addrChan <- addr
		}
		close(addrChan)
	}()

	return addrChan
}
