package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var ConfigLimiter limiter.Config = limiter.Config{
	Max:        10,
	Expiration: 60 * time.Second,
	LimitReached: func(c *fiber.Ctx) error {
		if c.Get("content-type") == "application/json" {
			return c.Status(418).JSON(fiber.Map{
				"msg":     "Too many requests, slow down I'm a teapot (10 requests per minute)",
				"success": false,
			})
		}
		return c.Render("views/home", fiber.Map{
			"ERR":                 "You have reached the limit of requests! Please check back later. (1 minute)",
			"CI_COMMIT_SHORT_SHA": CI_COMMIT_SHORT_SHA,
			"CI_COMMIT_BRANCH":    CI_COMMIT_BRANCH,
			"CI_BUILD":            CI_BUILD,
		})
	},
}
