package pg

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

type Util interface {
	DB() *pg.DB
}

type util struct {
	pg *pg.DB
}

func (u util) DB() *pg.DB {
	return u.pg
}

func Init() Util {
	var DB *pg.DB

	once.Do(func() {
		DB = pg.Connect(&pg.Options{
			User:     viper.GetString("postgres.username"),
			Password: viper.GetString("postgres.password"),
			Database: viper.GetString("postgres.database"),
			Addr:     viper.GetString("postgres.host") + ":" + viper.GetString("postgres.port"),
			OnConnect: func(_ context.Context, conn *pg.Conn) error {
				_, err := conn.Exec("set search_path=?", viper.GetString("postgres.schema"))
				if err != nil {
					panic(err.Error())
				}
				return nil
			},
		})
		if viper.GetString("profile") == "debug" {
			DB.AddQueryHook(dbLogger{})
		}

	})
	return &util{
		pg: DB,
	}
}
