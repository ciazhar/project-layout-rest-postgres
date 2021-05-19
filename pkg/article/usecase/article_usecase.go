package usecase

import (
	"github.com/ciazhar/project-layout-rest-postgres/pkg/article/model"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/article/repository/postgres"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/response"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/validator"
	"github.com/imdario/mergo"
	"time"
)

type ArticleUseCase interface {
	Fetch(param model.FetchParam) ([]model.FetchResponse, error)
	GetByID(id string) (model.Article, error)
	Store(req *model.Article) error
	Update(req *model.Article) error
	Delete(id string) error
}

type articleUseCase struct {
	validator         validator.Util
	ArticleRepository postgres.ArticlePostgresRepository
}

func (c articleUseCase) GetByID(id string) (model.Article, error) {
	return c.ArticleRepository.GetByID(id)
}

func (c articleUseCase) Update(req *model.Article) error {
	oldReq, err := c.ArticleRepository.GetByID(req.Id.String())
	if err != nil {
		return err
	}

	if err := mergo.Merge(req, oldReq); err != nil {
		return response.Error(err)
	}
	if err := c.validator.Struct(req); err != nil {
		return response.Error(err)
	}

	req.CreatedAt = oldReq.CreatedAt
	req.UpdatedAt = time.Now()
	req.DeletedAt = oldReq.DeletedAt

	return c.ArticleRepository.Update(req)
}

func (c articleUseCase) Delete(id string) error {
	return c.ArticleRepository.Delete(id)
}

func (c articleUseCase) Fetch(param model.FetchParam) ([]model.FetchResponse, error) {
	return c.ArticleRepository.Fetch(param)
}

func (c articleUseCase) Store(req *model.Article) error {
	if err := c.validator.Struct(req); err != nil {
		return response.Error(err)
	}
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
	return c.ArticleRepository.Store(req)
}

func NewArticleUseCase(validator validator.Util, ArticleRepository postgres.ArticlePostgresRepository) ArticleUseCase {
	return articleUseCase{
		validator:         validator,
		ArticleRepository: ArticleRepository,
	}
}
