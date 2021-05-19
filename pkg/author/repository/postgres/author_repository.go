package postgres

import (
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/model"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/query"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/query/pg"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/response"
	"github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"sync"
)

var once sync.Once

type AuthorPostgresRepository interface {
	Fetch(param model.FetchParam) ([]model.Author, error)
	GetByID(id string) (model.Author, error)
	Store(req *model.Author) error
	Update(req *model.Author) error
	Delete(id string) error
}

type repository struct {
	pg pg.Util
}

func (r repository) Delete(id string) error {
	object := &model.Author{
		Id: uuid.MustParse(id),
	}
	_, err := r.pg.DB().Model(object).WherePK().Delete(object)
	return response.Error(err)
}

func (r repository) Fetch(param model.FetchParam) ([]model.Author, error) {
	Authors := make([]model.Author, 0)
	offset, limit := query.ToOffsetLimit(param.Paginate, param.Page, param.Size)

	q := "SELECT * FROM fetch_author(?,?,?)"
	_, err := r.pg.DB().Query(&Authors, q, param.Name, limit, offset)
	return Authors, response.Error(err)
}

func (r repository) GetByID(id string) (model.Author, error) {
	author := model.Author{Id: uuid.MustParse(id)}
	if err := r.pg.DB().Model(&author).WherePK().Select(); err != nil {
		return author, response.Error(err)
	}
	return author, nil
}

func (r repository) Store(req *model.Author) error {
	req.Id = uuid.New()
	_, err := r.pg.DB().Model(req).Insert()
	return response.Error(err)
}

func (r repository) Update(req *model.Author) error {
	_, err := r.pg.DB().Model(req).WherePK().Update()
	return response.Error(err)
}

func NewAuthorPostgresRepository(pg pg.Util) AuthorPostgresRepository {
	once.Do(func() {
		if err := pg.DB().Model((*model.Author)(nil)).CreateTable(&orm.CreateTableOptions{
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
