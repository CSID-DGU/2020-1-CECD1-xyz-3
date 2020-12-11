package routes

import (
	"net/http"

	"github.com/gofiber/fiber"

	"github.com/doogeun/gateway/pkg/grpc"
	"github.com/doogeun/gateway/pkg/routes/v1beta1"
)

func Apply(app *fiber.App, jobSvc grpc.Client) {
	v1beta1.Apply(app.Group("/v1beta1"), jobSvc)
	app.Get("/ping", func(c *fiber.Ctx) {
		c.Status(http.StatusNoContent)
	})
}
