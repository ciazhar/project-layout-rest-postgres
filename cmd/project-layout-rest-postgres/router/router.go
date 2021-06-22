package router

import (
	"github.com/ciazhar/project-layout-rest-postgres/cmd/project-layout-rest-postgres/app"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/article"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/root"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/logger"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func Init(a app.Application) error {

	// Default error handler
	var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
		// Default unknown error
		code := response.CodeUnknownError

		if e, ok := err.(*fiber.Error); ok {
			// Override status code if fiber.Error type
			code = e.Code
		} else {
			// log error
			logger.Error(err.Error())
		}

		// Set Content-Type: text/plain; charset=utf-8
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		// Return statuscode with error message
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    code,
			"message": err.Error(),
		})
	}

	//init fiber and middleware
	r := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorHandler,
	})
	r.Use(cors.New())
	r.Use(pprof.New())
	r.Use(recover.New())
	a.Router = r

	//init module -> diurutkan berdasarkan dependency
	root.Init(a)
	author.Init(a)
	article.Init(a)

	//route 404
	r.Use(func(c *fiber.Ctx) error {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "route not found",
		})
	})

	//run
	logger.Info("appplication start in port : " + viper.GetString("port"))
	return r.Listen(":" + viper.GetString("port"))
}
