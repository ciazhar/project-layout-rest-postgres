package postgres

import (
	"context"
	"fmt"
	"github.com/ciazhar/project-layout-rest-postgres/pkg/author/model"
	"github.com/ciazhar/project-layout-rest-postgres/test"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/db"
	"github.com/ciazhar/project-layout-rest-postgres/third_party/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var pool *pgxpool.Pool
var ID uuid.UUID

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")                // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../../../../configs") // optionally look for config in the working directory
	err := viper.ReadInConfig()                // Find and read the config file
	if err != nil {                            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	logger.Init()
	pool = db.Init()
	sql := "truncate author cascade"
	if _, err := pool.Exec(context.Background(), sql); err != nil {
		panic(err.Error())
	}
}

func NewActual() model.Author {
	var Author model.Author
	test.ToStruct("author/actual.1.golden", &Author)
	Author.CreatedAt = time.Now()
	Author.UpdatedAt = time.Now()
	return Author
}

func NewActual2() model.Author {
	var Author model.Author
	test.ToStruct("author/actual.2.golden", &Author)
	Author.CreatedAt = time.Now()
	Author.UpdatedAt = time.Now()
	return Author
}

func TestRepositoryStore(t *testing.T) {
	actual := NewActual()
	actual2 := NewActual2()
	repo := NewAuthorPostgresRepository(pool)
	testCases := []struct {
		name    string
		payload model.Author
		error   error
	}{
		{"insert-default-success", actual, nil},
		{"insert-different-payload-success", actual2, nil},
		{"insert-duplicate-error", actual, &fiber.Error{
			Code:    2,
			Message: "ERROR: duplicate key value violates unique constraint \"author_name_uindex\" (SQLSTATE 23505)",
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := repo.Store(&testCase.payload)
			assert.Equal(t, err, testCase.error)

			if testCase.name == "insert-default-success" {
				ID = testCase.payload.ID
			}
		})
	}
}

func TestRepositoryFetch(t *testing.T) {
	actual := NewActual()
	actual2 := NewActual2()
	actualise := make([]model.Author, 0)
	actualise = append(actualise, actual)
	actualise = append(actualise, actual2)
	actualise2 := make([]model.Author, 0)
	actualise2 = append(actualise2, actual)
	one := 1

	param1 := model.FetchParam{}
	param2 := model.FetchParam{
		Page: &one,
		Size: &one,
	}

	testCases := []struct {
		name    string
		param   model.FetchParam
		payload []model.Author
		error   error
	}{
		{"fetch", param1, actualise, nil},
		{"fetch-with-paginate", param2, actualise2, nil},
		{"fetch-without-paginate", param1, actualise, nil},
		{"fetch-with-param-name", param1, actualise, nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := NewAuthorPostgresRepository(pool)
			expected, err := repo.Fetch(testCase.param)

			assert.Equal(t, len(expected), len(testCase.payload))
			assert.Equal(t, err, testCase.error)
		})
	}
}

//
func TestRepositoryGetByID(t *testing.T) {
	repo := NewAuthorPostgresRepository(pool)

	fmt.Println(ID)

	t.Run("default", func(t *testing.T) {
		expected, err := repo.GetByID(ID.String())

		assert.NotNil(t, expected)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		expected, err := repo.GetByID(uuid.New().String())

		assert.NotNil(t, expected)
		assert.Error(t, err)
	})
}

func TestRepositoryUpdate(t *testing.T) {
	actual := NewActual()
	repo := NewAuthorPostgresRepository(pool)

	t.Run("default", func(t *testing.T) {
		actual.ID = ID
		err := repo.Update(&actual)
		assert.NoError(t, err)
	})
}

func TestRepository_Delete(t *testing.T) {
	repo := NewAuthorPostgresRepository(pool)

	err := repo.Delete(ID.String())

	assert.NoError(t, err)
}
