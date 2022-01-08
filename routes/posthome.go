package routes

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"git.ucode.space/Phil/goshorly/db"
	"git.ucode.space/Phil/goshorly/utils"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type eurl struct {
	CLI bool   `json:"cli"`
	URL string `json:"surl" form:"surl"`
}

func Posthome(c *fiber.Ctx) error {
	u := new(eurl)
	if err := c.BodyParser(u); err != nil {

		if u.CLI {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"msg":     "Parsing Error",
			})
		}

		return c.Status(500).Render("views/home", fiber.Map{
			"ERR":            "Parsing Error",
			"GitCommitShort": utils.GitCommitShort,
			"GitBranch":      utils.GitBranch,
			"GitBuild":       utils.GitBuild,
		})
	}

	if !regexp.MustCompile(`^(http|https|mailto|ts3server)://`).MatchString(u.URL) {

		if u.CLI {
			return c.Status(424).JSON(&fiber.Map{
				"success": false,
				"msg":     "Invalid URL",
			})
		}

		return c.Status(424).Render("views/home", fiber.Map{
			"ERR":            "Invalid URL, please check and try again.",
			"GitCommitShort": utils.GitCommitShort,
			"GitBranch":      utils.GitBranch,
			"GitBuild":       utils.GitBuild,
		})
	}

	id, err := gonanoid.New(8)

	if err != nil {

		if u.CLI {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"msg":     err.Error(),
			})
		}

		return c.Status(500).Render("views/home", fiber.Map{
			"ERR":            err.Error(),
			"GitCommitShort": utils.GitCommitShort,
			"GitBranch":      utils.GitBranch,
			"GitBuild":       utils.GitBuild,
		})
	}

	err = db.Client.Set(id, u.URL, 1296000*time.Second).Err()

	if err != nil {
		if u.CLI {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"msg":     err.Error(),
			})
		}

		return c.Status(500).Render("views/home", fiber.Map{
			"ERR":            err.Error(),
			"GitCommitShort": utils.GitCommitShort,
			"GitBranch":      utils.GitBranch,
			"GitBuild":       utils.GitBuild,
		})
	}

	fURL := utils.URL + id

	val, err := db.Client.Get("created-links").Result()

	if err != nil {
		fmt.Println("Key not found, creating new one")
		err2 := db.Client.Set("created-links", "0", 0).Err()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
		fmt.Println("Set it the first time")
	} else {
		i, _ := strconv.Atoi(val)
		i++
		err3 := db.Client.Set("created-links", i, 0).Err()
		if err3 != nil {
			fmt.Println(err3.Error())
		} else {
			fmt.Println("New created value set to " + strconv.Itoa(i))
		}
	}

	if u.CLI {
		return c.Status(201).JSON(&fiber.Map{
			"success": true,
			"URL":     fURL,
		})
	}

	return c.Status(201).Render("views/home", fiber.Map{
		"URL":            fURL,
		"GitCommitShort": utils.GitCommitShort,
		"GitBranch":      utils.GitBranch,
		"GitBuild":       utils.GitBuild,
	})
}
