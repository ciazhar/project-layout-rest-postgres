package postgres

//func init() {
//	env.InitPath(env.GetEnvPath() + "/config.json")
//	query := query.PG()
//	query.DropTable((*model.Author)(nil), &orm.DropTableOptions{
//		IfExists: true,
//		Cascade:  true,
//	})
//	query.CreateTable((*model.Author)(nil), nil)
//}
//
//var ID string
//
//func NewActual() model.Author {
//	var Author model.Author
//	testdata.ToStruct("author/actual.1.golden", &Author)
//	Author.CreatedAt = time.Now()
//	Author.UpdatedAt = time.Now()
//	return Author
//}
//
//func NewActual2() model.Author {
//	var Author model.Author
//	testdata.ToStruct("author/actual.2.golden", &Author)
//	Author.CreatedAt = time.Now()
//	Author.UpdatedAt = time.Now()
//	return Author
//}
//
//func TestRepositoryStore(t *testing.T) {
//	actual := NewActual()
//	actual2 := NewActual2()
//	repo := NewAuthorPostgresRepository()
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
//		repo := NewAuthorPostgresRepository()
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
//		repo := NewAuthorPostgresRepository()
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
//	repo := NewAuthorPostgresRepository()
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
//	repo := NewAuthorPostgresRepository()
//
//	t.Run("default", func(t *testing.T) {
//		actual.Id = ID
//		err := repo.Update(&actual)
//		assert.NoError(t, err)
//	})
//}
