package routes

import (
	"time"

	"git.ucode.space/Phil/goshorly/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var ConfigLimiter limiter.Config = limiter.Config{
	Max:        10,
	Expiration: 60 * time.Second,
	LimitReached: func(c *fiber.Ctx) error {
		return c.Render("views/home", fiber.Map{
			"ERR":            "You have reached the limit of requests! Please check back later. (1 minute)",
			"GitCommitShort": utils.GitCommitShort,
			"GitBranch":      utils.GitBranch,
			"GitBuild":       utils.GitBuild,
		})
	},
}
