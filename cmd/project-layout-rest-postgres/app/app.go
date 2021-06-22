package app

import (
	"github.com/ciazhar/project-layout-rest-postgres/third_party/env"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/logger"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/query/pg"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/validator"
	"github.com/gofiber/fiber/v2"
	"os"
)

type Application struct {
	Validator validator.Util
	Postgres  pg.Util
	Router    fiber.Router
}

func Init(e string) (Application, error) {

	//init
	env.Init(e)
	logger.Init()
	v := validator.Init()
	p := pg.Init()

	//set default timezone
	if err := os.Setenv("TZ", "Asia/Jakarta"); err != nil {
		panic(err.Error())
	}

	return Application{
		Validator: v,
		Postgres:  p,
	}, nil
}
