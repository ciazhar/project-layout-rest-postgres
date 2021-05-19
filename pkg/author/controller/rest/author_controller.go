package rest

import (
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/model"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/usecase"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/response"
	"github.com/gofiber/fiber/v2"
)

type AuthorController interface {
	Fetch(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type authorController struct {
	AuthorUseCase usecase.AuthorUseCase
}

func (it authorController) Fetch(c *fiber.Ctx) error {
	param := model.FetchParam{}
	if err := c.QueryParser(&param); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	payload, err := it.AuthorUseCase.Fetch(param)
	if err != nil {
		return err
	}

	return response.Success(c, payload)
}

func (it authorController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	payload, err := it.AuthorUseCase.GetByID(id)
	if err != nil {
		return err
	}

	return response.Success(c, payload)
}

func (it authorController) Store(c *fiber.Ctx) error {
	var payload model.Author
	if err := c.BodyParser(&payload); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	if err := it.AuthorUseCase.Store(&payload); err != nil {
		return err
	}

	return response.Success(c, payload)
}

func (it authorController) Update(c *fiber.Ctx) error {
	var payload model.Author
	if err := c.BodyParser(&payload); err != nil {
		return response.Error(err, response.CodeBadRequest)
	}

	if err := it.AuthorUseCase.Update(&payload); err != nil {
		return err
	}

	return response.Success(c, payload)
}

func (it authorController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := it.AuthorUseCase.Delete(id); err != nil {
		return err
	}

	return response.Success(c, nil)
}

func NewAuthorController(AuthorUseCase usecase.AuthorUseCase) AuthorController {
	return authorController{AuthorUseCase: AuthorUseCase}
}
