package main

import (
	"fmt"
	"net/netip"
	"os"
	"sync"

	"github.com/iAziz786/ods/db"
	"github.com/iAziz786/ods/ips"
	flag "github.com/spf13/pflag"
)

var (
	cidrRange = flag.StringP("cidr", "r", "0.0.0.0/0", "The CIDR range in which to run the port scanning. Defaults to 0.0.0.0/0")
	help      = flag.BoolP("help", "h", false, "Show help screen")
	dialect   = flag.StringP("dialect", "d", "", "The dialect to use to query mongodb. Default to `mongodb`")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, `
ods -- Open MongoDB Scanner

Usage:
ods [--] [query options]

Query Options:
-r, --cidr=0.0.0.0/0
-p, --db-port=27017
-d, --dialect=mongodb
-h, --help`)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *dialect == "" {
		flag.Usage()
		os.Exit(1)
	}

	var pinger db.Pinger = db.NewPinger(*dialect)

	if pinger == nil {
		panic(
			fmt.Sprintf(`
dialect %s not supported
Request new dialects here: https://github.com/iAziz786/ods/issues/new
Please include the details on how to check if the db gets connected successfully.
`, *dialect))
	}

	vulnerableIPs := make(chan netip.Addr)
	var vulWg sync.WaitGroup

	vulWg.Add(1)
	go func() {
		for addr := range ips.CidrRange(*cidrRange) {
			vulWg.Add(1)
			go func(addr netip.Addr) {
				if pinger.Ping(addr) {
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
