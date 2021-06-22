package main

import (
	"github.com/ciazhar/project-layout-rest-postgres/cmd/project-layout-rest-postgres/app"
	"github.com/ciazhar/project-layout-rest-postgres/cmd/project-layout-rest-postgres/router"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {

	ca := &cli.App{
		Name:    viper.GetString("name"),
		Version: viper.GetString("version"),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "env",
				Aliases: []string{"e"},
				Usage:   "Load configuration from `FILE`",
				Value:   "config",
			},
		},
		Action: func(c *cli.Context) error {
			env := c.String("env")
			viper.SetConfigName(env)
			a, err := app.Init(env)
			if err != nil {
				panic(err)
			}
			return router.Init(a)
		},
	}

	if err := ca.Run(os.Args); err != nil {
		panic(err.Error())
	}
}
