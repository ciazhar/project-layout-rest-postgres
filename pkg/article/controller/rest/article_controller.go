package rest

import (
	"github.com/ciazhar/project-layout-rest-postgres/pkg/article/model"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/article/usecase"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/response"
	"github.com/gofiber/fiber/v2"
)

type ArticleController interface {
	Fetch(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type articleController struct {
	ArticleUseCase usecase.ArticleUseCase
}

func (it articleController) Fetch(c *fiber.Ctx) error {
	param := model.FetchParam{}
	if err := c.QueryParser(&param); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	payload, err := it.ArticleUseCase.Fetch(param)
	if err != nil {
		return err
	}

	return response.Success(c, payload)
}

func (it articleController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	payload, err := it.ArticleUseCase.GetByID(id)
	if err != nil {
		return err
	}

	return response.Success(c, payload)
}

func (it articleController) Store(c *fiber.Ctx) error {
	var payload model.Article
	if err := c.BodyParser(&payload); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	if err := it.ArticleUseCase.Store(&payload); err != nil {
		return err
	}

	return response.Success(c, payload)
}

func (it articleController) Update(c *fiber.Ctx) error {
	var payload model.Article
	if err := c.BodyParser(&payload); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	if err := it.ArticleUseCase.Update(&payload); err != nil {
		return err
	}

	return response.Success(c, payload)
}

func (it articleController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := it.ArticleUseCase.Delete(id); err != nil {
		return err
	}

	return response.Success(c, nil)
}

func NewArticleController(ArticleUseCase usecase.ArticleUseCase) ArticleController {
	return articleController{ArticleUseCase: ArticleUseCase}
}
