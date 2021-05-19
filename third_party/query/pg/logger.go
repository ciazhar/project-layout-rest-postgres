package pg

import (
	"context"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/logger"
	"github.com/go-pg/pg/v10"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, _ *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(_ context.Context, q *pg.QueryEvent) error {
	sql, _ := q.FormattedQuery()
	logger.Info(string(sql))
	return nil
}
