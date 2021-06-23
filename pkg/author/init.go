package author

import (
	"github.com/ciazhar/project-layout-rest-postgres/cmd/project-layout-rest-postgres/app"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/controller/rest"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/repository/postgres"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/usecase"
	postgres2 "github.com/ciazhar/project-layout-rest-postgres/pkg/author/validator/postgres"
)

func Init(app app.Application) {
	repo := postgres.NewAuthorPostgresRepository(app.DB)
	uc := usecase.NewAuthorUseCase(app.Validator, repo)
	controller := rest.NewAuthorController(uc)
	postgres2.NewAuthorPostgresValidator(app.Validator, repo)

	r := app.Router.Group("/author")
	r.Get("/", controller.Fetch)
	r.Get("/:id", controller.GetByID)
	r.Post("/", controller.Store)
	r.Put("/", controller.Update)
	r.Delete("/:id", controller.Delete)
}
