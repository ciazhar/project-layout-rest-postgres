package postgres

import (
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/repository/postgres"
	validator2 "github.com/ciazhar/project-layout-rest-postgres/third_party/validator"
	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"
)

type AuthorPostgresValidator interface {
	AuthorMustExist(fl validator.FieldLevel) bool
}

type authorPostgresValidator struct {
	AuthorRepository postgres.AuthorPostgresRepository
}

func (r authorPostgresValidator) AuthorMustExist(fl validator.FieldLevel) bool {
	id := fl.Field().Interface().(uuid.UUID)
	return r.validateAuthorID(id)
}

func (r authorPostgresValidator) validateAuthorID(id uuid.UUID) bool {
	if id != uuid.Nil {
		if _, err := r.AuthorRepository.GetByID(id.String()); err != nil {
			return false
		}
	}
	return true
}

func NewAuthorPostgresValidator(validator validator2.Util, AuthorRepository postgres.AuthorPostgresRepository) {
	v := authorPostgresValidator{
		AuthorRepository: AuthorRepository,
	}
	if err := validator.Validate().RegisterValidation("authorMustExist", v.AuthorMustExist); err != nil {
		panic(err.Error())
	}
}
