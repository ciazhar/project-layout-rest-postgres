package article

import (
	"github.com/ciazhar/project-layout-rest-postgres/cmd/project-layout-rest-postgres/app"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/article/controller/rest"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/article/repository/postgres"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/article/usecase"
)

func Init(app app.Application) {
	repo := postgres.NewArticlePostgresRepository(app.Postgres)
	uc := usecase.NewArticleUseCase(app.Validator, repo)
	controller := rest.NewArticleController(uc)

	r := app.Router.Group("/article")
	r.Get("/", controller.Fetch)
	r.Get("/:id", controller.GetByID)
	r.Post("/", controller.Store)
	r.Put("/", controller.Update)
	r.Delete("/:id", controller.Delete)
}
