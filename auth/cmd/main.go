package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/doogeun/auth/pkg/config"
	"github.com/doogeun/auth/pkg/database"
	"github.com/doogeun/auth/pkg/server"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.New(cfg.DatabaseURI, cfg.DatabaseTimeout)
	if err != nil {
		log.Fatal(err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	s := server.NewServer(*cfg, db)
	if err := s.Serve(sigCh); err != nil {
		log.Fatal(err)
	}
}
