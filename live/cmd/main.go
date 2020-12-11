package main

import (
	"fmt"

	"github.com/doogeun/live/pkg/cache"
	"github.com/doogeun/live/pkg/config"
	"github.com/doogeun/live/pkg/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	c := cache.New(cfg.CacheHost, cfg.CachePort)
	r := route.New(c)
	r.Apply(app)
	app.Listen(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
}
