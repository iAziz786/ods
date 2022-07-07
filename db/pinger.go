package db

import (
	"net/netip"

	"github.com/iAziz786/ods/db/mongo"
)

type Pinger interface {
	Ping(netip.Addr) bool
}

func NewPinger(dialect string) Pinger {
	switch dialect {
	case "mongodb":
		fallthrough
	case "mongodb+srv":
		return mongo.NewMongoPinger(dialect)
	default:
		return nil
	}
}
