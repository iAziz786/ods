package mongo

import (
	"context"
	"net/netip"
	"time"

	flag "github.com/spf13/pflag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoPort    = flag.StringP("mongo-port", "p", "27017", "The ports on which the scan should be run")
	mongoDialect = flag.StringP("mongo-dialect", "d", "mongodb", "The dialect to use to query mongodb. Default to `mongodb`")
)

func Ping(addr netip.Addr) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Client().ApplyURI(*mongoDialect + "://" + addr.String() + ":" + *mongoPort)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return false
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)

	return err == nil
}
