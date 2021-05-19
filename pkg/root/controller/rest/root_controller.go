package rest

import (
	"github.com/ciazhar/project-layout-rest-postgres/third_party/env"
	"github.com/gofiber/fiber/v2"
)

type RootController interface {
	Root(c *fiber.Ctx) error
}

type rootRestController struct {
	env env.Util
}

func NewRootRestController(util env.Util) RootController {
	return &rootRestController{
		env: util,
	}
}

func (r rootRestController) Root(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "welcome to " + r.env.Get("name") + " : " + r.env.Get("version") + " : " + r.env.Get("profile")})
}
