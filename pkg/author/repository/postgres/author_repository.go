package postgres

import (
	"context"
	"fmt"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/model"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/db"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/response"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthorPostgresRepository interface {
	Fetch(param model.FetchParam) ([]model.Author, error)
	GetByID(id string) (model.Author, error)
	Store(req *model.Author) error
	Update(req *model.Author) error
	Delete(id string) error
}

type repository struct {
	*pgxpool.Pool
	tableName string
}

func (r repository) Delete(id string) error {
	sql := fmt.Sprintf("update %s set deleted_at = now() where id = $1", r.tableName)
	_, err := r.Exec(context.Background(), sql, id)
	return response.Error(err)
}

func (r repository) Fetch(param model.FetchParam) ([]model.Author, error) {
	resp := make([]model.Author, 0)
	offset, limit := db.ToOffsetLimit(param.Paginate, param.Page, param.Size)
	sql := "SELECT * FROM fetch_author($1,$2,$3)"
	err := pgxscan.Select(context.Background(), r.Pool, &resp, sql, param.Name, limit, offset)
	return resp, response.Error(err)
}

func (r repository) GetByID(id string) (model.Author, error) {
	resp := model.Author{Id: uuid.MustParse(id)}
	sql := fmt.Sprintf("select * from %s where id=$1", r.tableName)
	err := pgxscan.Get(context.Background(), r.Pool, &resp, sql, id)
	return resp, response.Error(err)
}

func (r repository) Store(req *model.Author) error {
	req.Id = uuid.New()
	sql := fmt.Sprintf("insert into %s(id,name,created_at,updated_at) values($1,$2,now(),now())", r.tableName)
	_, err := r.Exec(context.Background(), sql, req.Id, req.Name)
	return response.Error(err)
}

func (r repository) Update(req *model.Author) error {
	sql := fmt.Sprintf(`
		update %s
		set name       = (case when $2 = '' then name else $2 end),
			updated_at = now()
		where id = $1
	`, r.tableName)
	_, err := r.Exec(context.Background(), sql, req.Id, req.Name)
	return response.Error(err)
}

func NewAuthorPostgresRepository(pool *pgxpool.Pool) AuthorPostgresRepository {
	return repository{
		Pool:      pool,
		tableName: "author",
	}
}
