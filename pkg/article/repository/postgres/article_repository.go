package postgres

import (
	"context"
	"fmt"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/article/model"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/db"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/response"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ArticlePostgresRepository interface {
	Fetch(param model.FetchParam) ([]model.FetchResponse, error)
	GetByID(id string) (model.Article, error)
	Store(req *model.Article) error
	Update(req *model.Article) error
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

func (r repository) Fetch(param model.FetchParam) ([]model.FetchResponse, error) {
	resp := make([]model.FetchResponse, 0)
	offset, limit := db.ToOffsetLimit(param.Paginate, param.Page, param.Size)
	sql := "SELECT * FROM fetch_article($1,$2,$3,$4,$5,$6)"
	if err := pgxscan.Select(context.Background(), r.Pool, &resp, sql, param.Title, param.AuthorID, param.From, param.Until, limit, offset); err != nil {
		return resp, response.Error(err)
	}
	return resp, nil
}

func (r repository) GetByID(id string) (model.Article, error) {
	resp := model.Article{ID: uuid.MustParse(id)}
	sql := fmt.Sprintf("select * from %s where id=$1", r.tableName)
	if err := pgxscan.Get(context.Background(), r.Pool, &resp, sql, id); err != nil {
		return resp, response.Error(err)
	}
	return resp, nil
}

func (r repository) Store(req *model.Article) error {
	req.ID = uuid.New()
	sql := fmt.Sprintf("insert into %s(id,title,content,author_id,created_at,updated_at) values($1,$2,$3,$4,now(),now())", r.tableName)
	if _, err := r.Exec(context.Background(), sql, req.ID, req.Title, req.Content, req.AuthorID); err != nil {
		return response.Error(err)
	}
	return nil
}

func (r repository) Update(req *model.Article) error {
	sql := fmt.Sprintf(`
		update %s
		set title       = (case when $2 = '' then title else $2 end),
			content       = (case when $3 = '' then content else $3 end),
			author_id       = (case when $4 is null then author_id else $4 end),
			updated_at = now()
		where id = $1
	`, r.tableName)
	if _, err := r.Exec(context.Background(), sql, req.ID, req.Title, req.Content, req.AuthorID); err != nil {
		return response.Error(err)
	}
	return nil
}

func NewArticlePostgresRepository(pool *pgxpool.Pool) ArticlePostgresRepository {
	return repository{
		Pool:      pool,
		tableName: "article",
	}
}
