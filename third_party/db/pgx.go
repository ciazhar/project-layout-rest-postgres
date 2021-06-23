package db

import (
	"context"
	"fmt"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/logger"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func Init() *pgxpool.Pool {
	urlExample := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s",
		viper.GetString("postgres.username"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.host"),
		viper.GetString("postgres.port"),
		viper.GetString("postgres.database"),
		viper.GetString("postgres.schema"))

	c, err := pgxpool.ParseConfig(urlExample)
	if err != nil {
		panic(err.Error())
	}
	if viper.GetString("profile") == "debug" {
		c.ConnConfig.Logger = zapadapter.NewLogger(logger.Logger)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), c)
	if err != nil {
		panic(err.Error())
	}
	return conn
}
