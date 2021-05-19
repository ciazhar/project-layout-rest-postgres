package pg

import (
	"context"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/env"
	"github.com/go-pg/pg/v10"
	"sync"
)

var once sync.Once

type Util interface {
	DB() *pg.DB
}

type util struct {
	env env.Util
	pg  *pg.DB
}

func (u util) DB() *pg.DB {
	return u.pg
}

func Init(env env.Util) Util {
	var DB *pg.DB

	once.Do(func() {
		DB = pg.Connect(&pg.Options{
			User:     env.Get("postgres.username"),
			Password: env.Get("postgres.password"),
			Database: env.Get("postgres.database"),
			Addr:     env.Get("postgres.host") + ":" + env.Get("postgres.port"),
			OnConnect: func(_ context.Context, conn *pg.Conn) error {
				_, err := conn.Exec("set search_path=?", env.Get("postgres.schema"))
				if err != nil {
					panic(err.Error())
				}
				return nil
			},
		})
		if env.Get("profile") == "debug" {
			DB.AddQueryHook(dbLogger{})
		}

	})
	return &util{
		env: env,
		pg:  DB,
	}
}
