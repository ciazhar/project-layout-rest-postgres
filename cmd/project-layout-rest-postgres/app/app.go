package app

import (
	"github.com/ciazhar/project-layout-rest-postgres/third_party/db"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/env"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/logger"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type Application struct {
	Validator validator.Util
	DB        *pgxpool.Pool
	Router    fiber.Router
}

func Init(e string) (Application, error) {

	//init
	env.Init(e)
	logger.Init()
	v := validator.Init()
	dbx := db.Init()

	//set default timezone
	if err := os.Setenv("TZ", "Asia/Jakarta"); err != nil {
		panic(err.Error())
	}

	return Application{
		Validator: v,
		DB:        dbx,
	}, nil
}
