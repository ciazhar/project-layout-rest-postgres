package usecase

//func NewActual() model.Article {
//	var article model.Article
//	testdata.ToStruct("article/actual.1.golden", &article)
//	return article
//}
//
//func NewActual2() model.Article {
//	var article model.Article
//	testdata.ToStruct("article/actual.2.golden", &article)
//	return article
//}
//
//func TestArticleUseCase_Store(t *testing.T) {
//	repo := new(mocks.ArticlePostgresRepository)
//	uc := NewArticleUseCase(repo)
//	testCases := []struct {
//		name        string
//		article model.Article
//		returnError error
//	}{
//		{"default", NewActual(), nil},
//		{"default2", NewActual2(), nil},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			repo.On("Store", &testCase.article).Return(testCase.returnError)
//
//			err := uc.Store(&testCase.article)
//
//			assert.NoError(t, err)
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestArticleUseCase_Fetch(t *testing.T) {
//	repo := new(mocks.ArticlePostgresRepository)
//	uc := NewArticleUseCase(repo)
//	testCases := []struct {
//		name             string
//		offset           int
//		limit            int
//		returnArticle []model.Article
//		returnError      error
//	}{
//		{"default", 0, 10, []model.Article{NewActual(), NewActual2()}, nil},
//		{"default2", 0, 5, []model.Article{NewActual(), NewActual2()}, nil},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//
//			param := rest.NewParam()
//			param.Offset = 1
//			param.Limit = 10
//
//			repo.On("Fetch", param).Return(testCase.returnArticle, testCase.returnError)
//
//			expected, err := uc.Fetch(param)
//
//			assert.NotEmpty(t, expected)
//			assert.NoError(t, err)
//			assert.Len(t, expected, len(testCase.returnArticle))
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestArticleUseCase_GetByID(t *testing.T) {
//	repo := new(mocks.ArticlePostgresRepository)
//	uc := NewArticleUseCase(repo)
//	testCases := []struct {
//		name             string
//		id               string
//		returnArticle model.Article
//		returnError      error
//	}{
//		{"default", "1", NewActual(), nil},
//		{"default2", "2", NewActual(), nil},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			repo.On("GetByID", testCase.id).Return(testCase.returnArticle, testCase.returnError)
//
//			expected, err := uc.GetByID(testCase.id)
//
//			assert.NoError(t, err)
//			assert.NotNil(t, expected)
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestArticleUseCase_Update(t *testing.T) {
//	repo := new(mocks.ArticlePostgresRepository)
//	uc := NewArticleUseCase(repo)
//	actual := NewActual()
//	actual.Id = "100"
//	testCases := []struct {
//		name        string
//		article model.Article
//		returnError error
//	}{
//		{"default", NewActual(), nil},
//		{"default2", actual, errors.New("not found")},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			repo.On("GetByID", testCase.article.Id).Return(testCase.article, testCase.returnError)
//			repo.On("Update", &testCase.article).Return(testCase.returnError)
//
//			err := uc.Update(&testCase.article)
//
//			assert.Equal(t, err, testCase.returnError)
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestArticleUseCase_Delete(t *testing.T) {
//	actual := NewActual()
//	actual.Id = "100"
//	repo := new(mocks.ArticlePostgresRepository)
//	uc := NewArticleUseCase(repo)
//	testCases := []struct {
//		name        string
//		article model.Article
//		returnError error
//	}{
//		{"default", NewActual(), nil},
//		{"error", actual, errors.New("not found")},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//
//			wayback := time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
//			patch := monkey.Patch(time.Now, func() time.Time { return wayback })
//			defer patch.Unpatch()
//
//			repo.On("GetByID", testCase.article.Id).Return(testCase.article, testCase.returnError)
//
//			testCase.article.DeletedAt = time.Now()
//			repo.On("Update", &testCase.article).Return(testCase.returnError)
//
//			err := uc.Delete(testCase.article.Id)
//
//			assert.Equal(t, err, testCase.returnError)
//			//repo.AssertExpectations(t)
//		})
//	}
//}
