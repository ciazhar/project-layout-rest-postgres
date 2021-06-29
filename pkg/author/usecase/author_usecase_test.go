package usecase

import (
	"errors"
	"github.com/ciazhar/project-layout-rest-postgres/internal/mocks"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/model"
	"github.com/ciazhar/project-layout-rest-postgres/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewActual() model.Author {
	var author model.Author
	test.ToStruct("author/actual.1.golden", &author)
	return author
}

func NewActual2() model.Author {
	var author model.Author
	test.ToStruct("author/actual.2.golden", &author)
	return author
}

func TestAuthorUseCase_Store(t *testing.T) {
	repo := new(mocks.AuthorPostgresRepository)
	validator := new(mocks.Util)
	uc := NewAuthorUseCase(validator, repo)
	testCases := []struct {
		name        string
		author      model.Author
		returnError error
	}{
		{"default", NewActual(), nil},
		{"default2", NewActual2(), nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("Store", &testCase.author).Return(testCase.returnError)
			validator.On("Struct", &testCase.author).Return(testCase.returnError)
			err := uc.Store(&testCase.author)

			assert.NoError(t, err)
			repo.AssertExpectations(t)
		})
	}
}

func TestAuthorUseCase_Fetch(t *testing.T) {
	repo := new(mocks.AuthorPostgresRepository)
	validator := new(mocks.Util)
	uc := NewAuthorUseCase(validator, repo)
	testCases := []struct {
		name         string
		page         int
		size         int
		returnAuthor []model.Author
		returnError  error
	}{
		{"default", 1, 10, []model.Author{NewActual(), NewActual2()}, nil},
		{"default2", 1, 5, []model.Author{NewActual(), NewActual2()}, nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			param := model.FetchParam{}
			param.Page = &testCase.page
			param.Size = &testCase.size

			repo.On("Fetch", param).Return(testCase.returnAuthor, testCase.returnError)

			expected, err := uc.Fetch(param)

			assert.NotEmpty(t, expected)
			assert.NoError(t, err)
			assert.Len(t, expected, len(testCase.returnAuthor))
			repo.AssertExpectations(t)
		})
	}
}

func TestAuthorUseCase_GetByID(t *testing.T) {
	repo := new(mocks.AuthorPostgresRepository)
	validator := new(mocks.Util)
	uc := NewAuthorUseCase(validator, repo)
	testCases := []struct {
		name         string
		id           string
		returnAuthor model.Author
		returnError  error
	}{
		{"default", "1", NewActual(), nil},
		{"default2", "2", NewActual(), nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("GetByID", testCase.id).Return(testCase.returnAuthor, testCase.returnError)

			expected, err := uc.GetByID(testCase.id)

			assert.NoError(t, err)
			assert.NotNil(t, expected)
			repo.AssertExpectations(t)
		})
	}
}

func TestAuthorUseCase_Update(t *testing.T) {
	repo := new(mocks.AuthorPostgresRepository)
	validator := new(mocks.Util)
	uc := NewAuthorUseCase(validator, repo)
	actual := NewActual()
	actual.ID = uuid.New()
	actual2 := NewActual2()
	actual2.ID = actual.ID
	actual2.Name = actual.Name

	testCases := []struct {
		name           string
		author         model.Author
		validatorError error
		updateError    error
	}{
		{"default", actual, nil, nil},
		//{"update-unique", actual2, errors.New(""), errors.New("")},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			validator.On("Struct", &testCase.author).Return(testCase.validatorError)
			repo.On("Update", &testCase.author).Return(testCase.updateError)

			err := uc.Update(&testCase.author)

			assert.Equal(t, err, testCase.updateError)
			repo.AssertExpectations(t)
		})
	}
}

func TestAuthorUseCase_Delete(t *testing.T) {
	actual := NewActual()
	actual.ID = uuid.New()
	actual2 := NewActual2()
	repo := new(mocks.AuthorPostgresRepository)
	validator := new(mocks.Util)
	uc := NewAuthorUseCase(validator, repo)
	testCases := []struct {
		name        string
		author      model.Author
		returnError error
	}{
		{"default", actual, nil},
		{"error", actual2, errors.New("not found")},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("Delete", testCase.author.ID.String()).Return(testCase.returnError)

			err := uc.Delete(testCase.author.ID.String())

			assert.Equal(t, err, testCase.returnError)
		})
	}
}
