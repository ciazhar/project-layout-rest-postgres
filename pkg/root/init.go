package root

import (
	"github.com/ciazhar/project-layout-rest-postgres/cmd/project-layout-rest-postgres/app"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/root/controller/rest"
)

func Init(app app.Application) {

	controller := rest.NewRootRestController()

	r := app.Router.Group("/")
	r.Get("/", controller.Root)

}
