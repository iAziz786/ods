package mongo

import (
	"context"
	"net/netip"
	"time"

	flag "github.com/spf13/pflag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoPinger struct {
	mongoPort    string
	mongoDialect string
}

func NewMongoPinger(dialect string) MongoPinger {
	mongoPort := flag.StringP("db-port", "p", "27017", "The ports on which the scan should be run")

	return MongoPinger{
		mongoPort:    *mongoPort,
		mongoDialect: dialect,
	}
}

func (m MongoPinger) Ping(addr netip.Addr) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Client().ApplyURI(m.mongoDialect + "://" + addr.String() + ":" + m.mongoPort)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return false
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)

	return err == nil
}
