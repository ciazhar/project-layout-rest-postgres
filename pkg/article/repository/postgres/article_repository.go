package postgres

import (
	"github.com/ciazhar/project-layout-rest-postgres/pkg/article/model"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/query"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/query/pg"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/response"
	"github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"sync"
)

var once sync.Once

type ArticlePostgresRepository interface {
	Fetch(param model.FetchParam) ([]model.FetchResponse, error)
	GetByID(id string) (model.Article, error)
	Store(req *model.Article) error
	Update(req *model.Article) error
	Delete(id string) error
}

type repository struct {
	pg pg.Util
}

func (r repository) Delete(id string) error {
	object := &model.Article{
		Id: uuid.MustParse(id),
	}
	_, err := r.pg.DB().Model(object).WherePK().Delete(object)
	return response.Error(err)
}

func (r repository) Fetch(param model.FetchParam) ([]model.FetchResponse, error) {

	offset, limit := query.ToOffsetLimit(param.Paginate, param.Page, param.Size)

	resp := make([]model.FetchResponse, 0)
	q := "SELECT * FROM fetch_article(?,?,?,?,?,?)"
	_, err := r.pg.DB().Query(&resp, q, param.Title, param.AuthorID, param.From, param.Until, limit, offset)
	return resp, response.Error(err)
}

func (r repository) GetByID(id string) (model.Article, error) {
	Article := model.Article{Id: uuid.MustParse(id)}
	if err := r.pg.DB().Model(&Article).WherePK().Select(); err != nil {
		return Article, response.Error(err)
	}
	return Article, nil
}

func (r repository) Store(req *model.Article) error {
	req.Id = uuid.New()
	_, err := r.pg.DB().Model(req).Insert()
	return response.Error(err)
}

func (r repository) Update(req *model.Article) error {
	_, err := r.pg.DB().Model(req).WherePK().Update()
	return response.Error(err)
}

func NewArticlePostgresRepository(pg pg.Util) ArticlePostgresRepository {
	once.Do(func() {
		if err := pg.DB().Model((*model.Article)(nil)).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			Temp:          false,
			FKConstraints: true,
		}); err != nil {
			panic(err)
		}
	})
	return repository{
		pg: pg,
	}
}
