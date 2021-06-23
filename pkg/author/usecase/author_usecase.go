package usecase

import (
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/model"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/repository/postgres"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/response"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/validator"
)

type AuthorUseCase interface {
	Fetch(param model.FetchParam) ([]model.Author, error)
	GetByID(id string) (model.Author, error)
	Store(req *model.Author) error
	Update(req *model.Author) error
	Delete(id string) error
}

type authorUseCase struct {
	validator        validator.Util
	AuthorRepository postgres.AuthorPostgresRepository
}

func (c authorUseCase) GetByID(id string) (model.Author, error) {
	return c.AuthorRepository.GetByID(id)
}

func (c authorUseCase) Update(req *model.Author) error {
	if err := c.validator.Struct(req); err != nil {
		return response.Error(err)
	}
	return c.AuthorRepository.Update(req)
}

func (c authorUseCase) Delete(id string) error {
	return c.AuthorRepository.Delete(id)
}

func (c authorUseCase) Fetch(param model.FetchParam) ([]model.Author, error) {
	return c.AuthorRepository.Fetch(param)
}

func (c authorUseCase) Store(req *model.Author) error {
	if err := c.validator.Struct(req); err != nil {
		return response.Error(err)
	}
	return c.AuthorRepository.Store(req)
}

func NewAuthorUseCase(validator validator.Util, AuthorRepository postgres.AuthorPostgresRepository) AuthorUseCase {
	return authorUseCase{
		validator:        validator,
		AuthorRepository: AuthorRepository,
	}
}
