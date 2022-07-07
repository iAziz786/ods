package main

import (
	"fmt"
	"net/netip"
	"os"
	"sync"

	"github.com/iAziz786/oms/ips"
	"github.com/iAziz786/oms/mongo"
	flag "github.com/spf13/pflag"
)

var (
	cidrRange = flag.StringP("cidr", "r", "0.0.0.0/0", "The CIDR range in which to run the port scanning. Defaults to 0.0.0.0/0")
	help      = flag.BoolP("help", "h", false, "Show help screen")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, `
oms -- Open MongoDB Scanner

Usage:
oms [--] [query options]

Query Options:
-r, --cidr=0.0.0.0/0
-p, --mongo-port=27017
-d, --mongo-dialect=mongodb
-h, --help`)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	vulnerableIPs := make(chan netip.Addr)
	var vulWg sync.WaitGroup

	vulWg.Add(1)
	go func() {
		for addr := range ips.CidrRange(*cidrRange) {
			vulWg.Add(1)
			go func(addr netip.Addr) {
				if mongo.Ping(addr) {
					vulnerableIPs <- addr
				}
				vulWg.Done()
			}(addr)
		}
		vulWg.Done()
	}()

	go func() {
		vulWg.Wait()
		close(vulnerableIPs)
	}()

	for vulnerableIP := range vulnerableIPs {
		fmt.Println(vulnerableIP)
	}
}

func init() {
	flag.Parse()
}
