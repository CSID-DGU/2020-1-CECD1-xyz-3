package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/doogeun/live/pkg/auth"
	cachepkg "github.com/doogeun/live/pkg/cache"

	"github.com/gofiber/fiber/v2"
)

type Route interface {
	Apply(a *fiber.App)
}

type route struct {
	cache cachepkg.Cache
}

type tokenRequest struct {
	Email string `json:"email"`
}

type tokenResponse struct {
	AccessToken string `json:"accessToken"`
}

const (
	applicantsKey = "applicants"
	positionsKey  = "positions"
)

func New(cache cachepkg.Cache) Route {
	return &route{
		cache: cache,
	}
}

func (r *route) Apply(a *fiber.App) {
	a.Get("/v1/applicants", func(c *fiber.Ctx) error {
		return r.getJSONResponse(c, applicantsKey)
	})
	a.Put("/v1/applicants", func(c *fiber.Ctx) error {
		return r.setJSON(c, applicantsKey)
	})
	a.Get("/v1/positions", func(c *fiber.Ctx) error {
		return r.getJSONResponse(c, positionsKey)
	})
	a.Put("/v1/positions", func(c *fiber.Ctx) error {
		return r.setJSON(c, positionsKey)
	})
	a.Get("/v1/me", func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		key := fmt.Sprintf("token.%s", token)
		data, err := r.cache.GetJSON(c.Context(), key)
		if err != nil {
			return err
		}
		if data == nil {
			return c.SendStatus(http.StatusUnauthorized)
		}
		return c.JSON(data)
	})
	a.Post("/v1/token", func(c *fiber.Ctx) error {
		req := tokenRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		token, err := auth.NewToken(req.Email)
		if err != nil {
			return err
		}
		key := fmt.Sprintf("token.%s", token)
		if err := r.cache.SetJSON(c.Context(), key, c.Body(), 1*time.Hour); err != nil {
			return err
		}
		c.Status(http.StatusCreated)
		return c.JSON(tokenResponse{AccessToken: token})
	})
}

func (r *route) getJSONResponse(c *fiber.Ctx, key string) error {
	data, err := r.cache.GetJSON(c.Context(), key)
	if err != nil {
		return err
	}
	if data == nil {
		return c.SendStatus(http.StatusNotFound)
	}
	return c.JSON(data)
}

func (r *route) setJSON(c *fiber.Ctx, key string, ttl ...time.Duration) error {
	var msg json.RawMessage
	if err := json.Unmarshal(c.Body(), &msg); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	if err := r.cache.SetJSON(c.Context(), key, msg, ttl...); err != nil {
		return err
	}
	return c.SendStatus(http.StatusOK)
}
