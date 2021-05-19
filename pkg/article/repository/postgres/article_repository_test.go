package postgres

//func init() {
//	env.InitPath(env.GetEnvPath() + "/config.json")
//	query := query.PG()
//	query.DropTable((*model.Article)(nil), &orm.DropTableOptions{
//		IfExists: true,
//		Cascade:  true,
//	})
//	query.CreateTable((*model.Article)(nil), nil)
//}
//
//var ID string
//
//func NewActual() model.Article {
//	var Article model.Article
//	testdata.ToStruct("article/actual.1.golden", &Article)
//	Article.CreatedAt = time.Now()
//	Article.UpdatedAt = time.Now()
//	return Article
//}
//
//func NewActual2() model.Article {
//	var Article model.Article
//	testdata.ToStruct("article/actual.2.golden", &Article)
//	Article.CreatedAt = time.Now()
//	Article.UpdatedAt = time.Now()
//	return Article
//}
//
//func TestRepositoryStore(t *testing.T) {
//	actual := NewActual()
//	actual2 := NewActual2()
//	repo := NewArticlePostgresRepository()
//
//	t.Run("default", func(t *testing.T) {
//		err := repo.Store(&actual)
//		assert.NoError(t, err)
//		ID = actual.Id
//	})
//	t.Run("default2", func(t *testing.T) {
//		err := repo.Store(&actual2)
//		assert.NoError(t, err)
//	})
//}
//
//func TestRepositoryFetch(t *testing.T) {
//	t.Run("default", func(t *testing.T) {
//		repo := NewArticlePostgresRepository()
//		param := rest.NewParam()
//		param.Offset = 0
//		param.Limit = 10
//		expected, err := repo.Fetch(param)
//
//		assert.NotEmpty(t, expected)
//		assert.NoError(t, err)
//		assert.Len(t, expected, 2)
//	})
//
//	t.Run("error", func(t *testing.T) {
//		repo := NewArticlePostgresRepository()
//		param := rest.NewParam()
//		param.Offset = 0
//		param.Limit = -10
//		expected, err := repo.Fetch(param)
//
//		assert.Empty(t, expected)
//		assert.Error(t, err)
//	})
//}
//
//func TestRepositoryGetByID(t *testing.T) {
//	repo := NewArticlePostgresRepository()
//
//	t.Run("default", func(t *testing.T) {
//		expected, err := repo.GetByID(ID)
//
//		assert.NotNil(t, expected)
//		assert.NoError(t, err)
//	})
//
//	t.Run("error", func(t *testing.T) {
//		expected, err := repo.GetByID("100")
//
//		assert.NotNil(t, expected)
//		assert.Error(t, err)
//	})
//}
//
//func TestRepositoryUpdate(t *testing.T) {
//	actual := NewActual()
//	repo := NewArticlePostgresRepository()
//
//	t.Run("default", func(t *testing.T) {
//		actual.Id = ID
//		err := repo.Update(&actual)
//		assert.NoError(t, err)
//	})
//}
