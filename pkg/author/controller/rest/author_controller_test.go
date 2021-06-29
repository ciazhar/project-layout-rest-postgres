package rest

import (
	"github.com/ciazhar/project-layout-rest-postgres/internal/mocks"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/model"
	"github.com/ciazhar/project-layout-rest-postgres/test"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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

func NewActualReader(payload model.Author) *strings.Reader {
	return test.ToReader(payload)
}

func TestAuthorController_Fetch(t *testing.T) {
	testCases := []struct {
		name         string
		page         int
		size         int
		returnAuthor []model.Author
		returnError  error
		httpStatus   int
	}{
		{"default", 1, 10, []model.Author{NewActual(), NewActual2()}, nil, http.StatusOK},
		//{"error", 1, 5, []model.Author{NewActual(), NewActual2()}, errors.New(""), http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			uc := new(mocks.AuthorUseCase)
			ctrl := NewAuthorController(uc)
			param := model.FetchParam{
				Page: &testCase.page,
				Size: &testCase.size,
			}

			router := fiber.New()
			router.Get("/author", ctrl.Fetch)
			uc.On("Fetch", param).Return(testCase.returnAuthor, testCase.returnError)
			r, err := router.Test(httptest.NewRequest(http.MethodGet, "/author?page="+strconv.Itoa(testCase.page)+"&size="+strconv.Itoa(testCase.size), nil))
			assert.NoError(t, err)
			assert.Equal(t, testCase.httpStatus, r.StatusCode)
			uc.AssertExpectations(t)
		})
	}
}

func TestAuthorController_GetByID(t *testing.T) {
	testCases := []struct {
		name         string
		id           string
		returnAuthor model.Author
		returnError  error
		httpStatus   int
	}{
		{"default", "1", NewActual(), nil, http.StatusOK},
		//{"bad-request","bukan-id",model.Author{}, errors.New(""), http.StatusBadRequest},
		//{"internal-server-error", "10", model.Author{}, errors.New(""), http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			uc := new(mocks.AuthorUseCase)
			ctrl := NewAuthorController(uc)

			router := fiber.New()
			router.Get("/author/:id", ctrl.GetByID)
			uc.On("GetByID", testCase.id).Return(testCase.returnAuthor, testCase.returnError)
			r, err := router.Test(httptest.NewRequest(http.MethodGet, "/author/"+testCase.id, nil))
			assert.NoError(t, err)
			assert.Equal(t, testCase.httpStatus, r.StatusCode)
			uc.AssertExpectations(t)
		})
	}
}

func TestAuthorController_Store(t *testing.T) {
	testCases := []struct {
		name        string
		payload     model.Author
		reader      io.Reader
		returnError error
		httpStatus  int
	}{
		{"default", NewActual(), NewActualReader(NewActual()), nil, http.StatusOK},
		//{"bad-request", NewActual(), strings.NewReader("{"), errors.New(""), http.StatusBadRequest},
		//{"error", NewActual2(), NewActualReader(NewActual2()), errors.New(""), http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			uc := new(mocks.AuthorUseCase)
			ctrl := NewAuthorController(uc)

			router := fiber.New()
			router.Post("/author", ctrl.Store)
			uc.On("Store", &testCase.payload).Return(testCase.returnError)
			r, err := router.Test(httptest.NewRequest(http.MethodPost, "/author", testCase.reader))
			assert.NoError(t, err)
			assert.Equal(t, testCase.httpStatus, r.StatusCode)
			uc.AssertExpectations(t)
		})
	}
}

func TestAuthorController_Update(t *testing.T) {
	testCases := []struct {
		name        string
		payload     model.Author
		reader      io.Reader
		returnError error
		httpStatus  int
	}{
		{"default", NewActual(), NewActualReader(NewActual()), nil, http.StatusOK},
		//{"bad-request", NewActual(), strings.NewReader("{"), errors.New(""), http.StatusBadRequest},
		//{"error", NewActual2(), NewActualReader(NewActual2()), errors.New(""), http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			uc := new(mocks.AuthorUseCase)
			ctrl := NewAuthorController(uc)

			router := fiber.New()
			router.Put("/author", ctrl.Update)
			uc.On("Update", &testCase.payload).Return(testCase.returnError)
			r, err := router.Test(httptest.NewRequest(http.MethodPut, "/author", testCase.reader))
			assert.NoError(t, err)
			assert.Equal(t, testCase.httpStatus, r.StatusCode)
			uc.AssertExpectations(t)
		})
	}
}

func TestAuthorController_Delete(t *testing.T) {
	testCases := []struct {
		name        string
		id          string
		returnError error
		httpStatus  int
	}{
		{"default", "1", nil, http.StatusOK},
		//{"error", "10", errors.New(""), http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			uc := new(mocks.AuthorUseCase)
			ctrl := NewAuthorController(uc)

			router := fiber.New()
			router.Delete("/author/:id", ctrl.Delete)
			uc.On("Delete", testCase.id).Return(testCase.returnError)
			r, err := router.Test(httptest.NewRequest(http.MethodDelete, "/author/"+testCase.id, nil))
			assert.NoError(t, err)
			assert.Equal(t, testCase.httpStatus, r.StatusCode)
			uc.AssertExpectations(t)
		})
	}
}
