package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"

	"github.com/doogeun/gateway/pkg/config"
	"github.com/doogeun/gateway/pkg/grpc"
	"github.com/doogeun/gateway/pkg/routes"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	jobSvc, err := grpc.NewClient(
		cfg.JobServiceURI, cfg.JobServiceAudience, cfg.SecureServiceConnection, cfg.SecureGCPServiceConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer jobSvc.Close()
	app := fiber.New()
	routes.Apply(app, jobSvc)
	if err := app.Listen(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)); err != nil {
		log.Fatal(err)
	}
}
