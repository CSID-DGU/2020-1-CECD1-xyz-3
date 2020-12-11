package server

import (
	"context"
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/doogeun/auth/pkg/config"
	"github.com/doogeun/auth/pkg/database"
	"github.com/doogeun/auth/pkg/repositories/companies"
	"github.com/doogeun/auth/pkg/server/healthcheck"
	"github.com/doogeun/auth/pkg/server/v1beta1"
	authpbv1b1 "github.com/doogeun/idl/gen/go/auth/v1beta1"
)

type Server struct {
	server   *grpc.Server
	db       database.Client
	cfg      config.Config
	cancelFn context.CancelFunc
	statusCh chan<- healthcheck.Status
}

func NewServer(c config.Config, db database.Client) *Server {
	s := grpc.NewServer()
	ctx, cancel := context.WithCancel(context.Background())
	healthChecker, sc := healthcheck.NewChecker(ctx)
	grpc_health_v1.RegisterHealthServer(s, healthChecker)
	authpbv1b1.RegisterAPIServiceServer(s, v1beta1.New(companies.New(db)))
	return &Server{
		server:   s,
		db:       db,
		cfg:      c,
		cancelFn: cancel,
		statusCh: sc,
	}
}

func (s *Server) Serve(c <-chan os.Signal) error {
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	serveErrCh := make(chan error)
	go func() {
		log.Infof("server listen on %s", addr)
		if err := s.server.Serve(l); err != nil {
			serveErrCh <- err
		}
	}()
	select {
	case _ = <-c:
		s.Stop()
	case err := <-serveErrCh:
		return err
	}
	return nil
}

func (s *Server) Stop() {
	defer s.cancelFn()
	s.statusCh <- healthcheck.Unhealthy
	s.server.GracefulStop()
}
