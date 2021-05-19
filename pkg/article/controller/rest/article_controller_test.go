package rest

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
//func NewActualReader(payload model.Article) *strings.Reader {
//	return testdata.ToReader(payload)
//}
//
//func TestArticleController_Fetch(t *testing.T) {
//	testCases := []struct {
//		name             string
//		offset           int
//		limit            int
//		returnArticle []model.Article
//		returnError      error
//		httpStatus       int
//	}{
//		{"default", 0, 10, []model.Article{NewActual(), NewActual2()}, nil, http.StatusOK},
//		{"error", 1, 5, []model.Article{NewActual(), NewActual2()}, errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.ArticleUseCase)
//			ctrl := NewArticleController(uc)
//
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.GET("/article", ctrl.Fetch)
//
//			r, err := http.NewRequest(http.MethodGet, "/article?offset="+strconv.Itoa(testCase.offset)+"&limit="+strconv.Itoa(testCase.limit), nil)
//			assert.NoError(t, err)
//			param := rest.NewParam()
//			param.Limit = testCase.limit
//			param.Offset = testCase.offset
//			uc.On("Fetch", param).Return(testCase.returnArticle, testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//			uc.AssertExpectations(t)
//		})
//	}
//
//}
//
//func TestArticleController_GetByID(t *testing.T) {
//	testCases := []struct {
//		name             string
//		id               string
//		returnArticle model.Article
//		returnError      error
//		httpStatus       int
//	}{
//		{"default", "1", NewActual(), nil, http.StatusOK},
//		//{"bad-request","bukan-id",model.Article{}, errors.New(""), http.StatusBadRequest},
//		{"internal-server-error", "10", model.Article{}, errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.ArticleUseCase)
//			ctrl := NewArticleController(uc)
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.GET("/article/:id", ctrl.GetByID)
//			r, err := http.NewRequest(http.MethodGet, "/article/"+testCase.id, nil)
//			assert.NoError(t, err)
//			uc.On("GetByID", testCase.id).Return(testCase.returnArticle, testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//			uc.AssertExpectations(t)
//		})
//	}
//}
//
//func TestArticleController_Store(t *testing.T) {
//	testCases := []struct {
//		name        string
//		payload     model.Article
//		reader      io.Reader
//		returnError error
//		httpStatus  int
//	}{
//		{"default", NewActual(), NewActualReader(NewActual()), nil, http.StatusOK},
//		{"bad-request", NewActual(), strings.NewReader("{"), errors.New(""), http.StatusBadRequest},
//		{"error", NewActual2(), NewActualReader(NewActual2()), errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//
//			uc := new(mocks.ArticleUseCase)
//			ctrl := NewArticleController(uc)
//			w := httptest.NewRecorder()
//			router := gin.New()
//			router.POST("/article", ctrl.Store)
//
//			r, err := http.NewRequest(http.MethodPost, "/article", testCase.reader)
//			assert.NoError(t, err)
//			uc.On("Store", &testCase.payload).Return(testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//		})
//	}
//}
//
//func TestArticleController_Update(t *testing.T) {
//	testCases := []struct {
//		name        string
//		payload     model.Article
//		reader      io.Reader
//		returnError error
//		httpStatus  int
//	}{
//		{"default", NewActual(), NewActualReader(NewActual()), nil, http.StatusOK},
//		{"bad-request", NewActual(), strings.NewReader("{"), errors.New(""), http.StatusBadRequest},
//		{"error", NewActual2(), NewActualReader(NewActual2()), errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.ArticleUseCase)
//			ctrl := NewArticleController(uc)
//			w := httptest.NewRecorder()
//			router := gin.New()
//			router.PUT("/article", ctrl.Update)
//
//			r, err := http.NewRequest(http.MethodPut, "/article", testCase.reader)
//			assert.NoError(t, err)
//			uc.On("Update", &testCase.payload).Return(testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//		})
//	}
//}
//
//func TestArticleController_Delete(t *testing.T) {
//	testCases := []struct {
//		name        string
//		id          string
//		returnError error
//		httpStatus  int
//	}{
//		{"default", "1", nil, http.StatusOK},
//		{"error", "10", errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.ArticleUseCase)
//			ctrl := NewArticleController(uc)
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.DELETE("/article/:id", ctrl.Delete)
//
//			r, err := http.NewRequest(http.MethodDelete, "/article/"+testCase.id, nil)
//			assert.NoError(t, err)
//			uc.On("Delete", testCase.id).Return(testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//			uc.AssertExpectations(t)
//		})
//	}
//}
