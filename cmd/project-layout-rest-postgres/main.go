package main

import (
	"github.com/ciazhar/project-layout-rest-postgres/cmd/project-layout-rest-postgres/app"
	"github.com/ciazhar/project-layout-rest-postgres/cmd/project-layout-rest-postgres/router"
	"github.com/urfave/cli"
	"os"
)

func main() {

	//init app
	a, err := app.Init()
	if err != nil {
		panic(err)
	}
	clientApp := cli.NewApp()
	clientApp.Name = a.Env.Get("name")
	clientApp.Version = a.Env.Get("version")
	clientApp.HideVersion = true
	clientApp.HideHelp = true
	clientApp.Action = router.Init(a)
	err = clientApp.Run(os.Args)
	panic(err)
}
