package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	raw         *mongo.Client
	connTimeout time.Duration
}

type Client interface {
	Database(name string) *mongo.Database
	Teardown() error
}

func New(uri string, connTimeout time.Duration) (Client, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()
	if err := c.Connect(ctx); err != nil {
		return nil, err
	}
	return &client{
		raw:         c,
		connTimeout: connTimeout,
	}, nil
}

func (c *client) Database(name string) *mongo.Database {
	return c.raw.Database(name, options.Database())
}

func (c *client) Teardown() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.connTimeout)
	defer cancel()
	return c.raw.Disconnect(ctx)
}

func Limit(limit int32) *int64 {
	v := int64(limit)
	return &v
}
