package main

import (
	"github.com/ciazhar/project-layout-rest-postgres/cmd/project-layout-rest-postgres/app"
	"github.com/ciazhar/project-layout-rest-postgres/cmd/project-layout-rest-postgres/router"
)

func main() {

	//init app
	a, err := app.Init()
	if err != nil {
		panic(err)
	}

	//init router
	if err := router.Init(a); err != nil {
		panic(err)
	}

}
