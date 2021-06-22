package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type RootController interface {
	Root(c *fiber.Ctx) error
}

type rootRestController struct{}

func NewRootRestController() RootController {
	return &rootRestController{}
}

func (r rootRestController) Root(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "welcome to " + viper.GetString("name") + " : " + viper.GetString("version") + " : " + viper.GetString("profile")})
}
