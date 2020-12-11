package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

var tokenSourceTimeout = 30 * time.Second

type client struct {
	conn        *grpc.ClientConn
	tokenSource oauth2.TokenSource
}

type Client interface {
	Connection() grpc.ClientConnInterface
	Context(parent context.Context) (context.Context, error)
	Close() error
}

func NewClient(uri, audience string, secure, gcpSecure bool) (Client, error) {
	conn, err := newConnection(uri, secure)
	if err != nil {
		return nil, err
	}
	var tokenSource oauth2.TokenSource
	if gcpSecure {
		tokenSource, err = newTokenSource(audience)
		if err != nil {
			return nil, err
		}
	}
	return &client{
		conn:        conn,
		tokenSource: tokenSource,
	}, nil
}

func newConnection(uri string, secure bool) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithAuthority(uri)}
	if secure {
		certPool, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{RootCAs: certPool})))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	return grpc.Dial(uri, opts...)
}

func newTokenSource(audience string) (oauth2.TokenSource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), tokenSourceTimeout)
	defer cancel()
	return idtoken.NewTokenSource(ctx, audience)
}

func (c *client) Connection() grpc.ClientConnInterface {
	return c.conn
}

func (c *client) Context(parent context.Context) (context.Context, error) {
	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, err
	}
	return metadata.AppendToOutgoingContext(
		parent, "authorization", fmt.Sprintf("Bearer %s", token.AccessToken)), nil
}

func (c *client) Close() error {
	return c.conn.Close()
}
