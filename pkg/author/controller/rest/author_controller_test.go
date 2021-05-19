package rest

//func NewActual() model.Author {
//	var author model.Author
//	testdata.ToStruct("author/actual.1.golden", &author)
//	return author
//}
//
//func NewActual2() model.Author {
//	var author model.Author
//	testdata.ToStruct("author/actual.2.golden", &author)
//	return author
//}
//
//func NewActualReader(payload model.Author) *strings.Reader {
//	return testdata.ToReader(payload)
//}
//
//func TestAuthorController_Fetch(t *testing.T) {
//	testCases := []struct {
//		name             string
//		offset           int
//		limit            int
//		returnAuthor []model.Author
//		returnError      error
//		httpStatus       int
//	}{
//		{"default", 0, 10, []model.Author{NewActual(), NewActual2()}, nil, http.StatusOK},
//		{"error", 1, 5, []model.Author{NewActual(), NewActual2()}, errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.AuthorUseCase)
//			ctrl := NewAuthorController(uc)
//
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.GET("/author", ctrl.Fetch)
//
//			r, err := http.NewRequest(http.MethodGet, "/author?offset="+strconv.Itoa(testCase.offset)+"&limit="+strconv.Itoa(testCase.limit), nil)
//			assert.NoError(t, err)
//			param := rest.NewParam()
//			param.Limit = testCase.limit
//			param.Offset = testCase.offset
//			uc.On("Fetch", param).Return(testCase.returnAuthor, testCase.returnError)
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
//func TestAuthorController_GetByID(t *testing.T) {
//	testCases := []struct {
//		name             string
//		id               string
//		returnAuthor model.Author
//		returnError      error
//		httpStatus       int
//	}{
//		{"default", "1", NewActual(), nil, http.StatusOK},
//		//{"bad-request","bukan-id",model.Author{}, errors.New(""), http.StatusBadRequest},
//		{"internal-server-error", "10", model.Author{}, errors.New(""), http.StatusInternalServerError},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			uc := new(mocks.AuthorUseCase)
//			ctrl := NewAuthorController(uc)
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.GET("/author/:id", ctrl.GetByID)
//			r, err := http.NewRequest(http.MethodGet, "/author/"+testCase.id, nil)
//			assert.NoError(t, err)
//			uc.On("GetByID", testCase.id).Return(testCase.returnAuthor, testCase.returnError)
//
//			router.ServeHTTP(w, r)
//
//			assert.Equal(t, testCase.httpStatus, w.Code)
//			uc.AssertExpectations(t)
//		})
//	}
//}
//
//func TestAuthorController_Store(t *testing.T) {
//	testCases := []struct {
//		name        string
//		payload     model.Author
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
//			uc := new(mocks.AuthorUseCase)
//			ctrl := NewAuthorController(uc)
//			w := httptest.NewRecorder()
//			router := gin.New()
//			router.POST("/author", ctrl.Store)
//
//			r, err := http.NewRequest(http.MethodPost, "/author", testCase.reader)
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
//func TestAuthorController_Update(t *testing.T) {
//	testCases := []struct {
//		name        string
//		payload     model.Author
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
//			uc := new(mocks.AuthorUseCase)
//			ctrl := NewAuthorController(uc)
//			w := httptest.NewRecorder()
//			router := gin.New()
//			router.PUT("/author", ctrl.Update)
//
//			r, err := http.NewRequest(http.MethodPut, "/author", testCase.reader)
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
//func TestAuthorController_Delete(t *testing.T) {
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
//			uc := new(mocks.AuthorUseCase)
//			ctrl := NewAuthorController(uc)
//			w := httptest.NewRecorder()
//
//			router := gin.New()
//			router.DELETE("/author/:id", ctrl.Delete)
//
//			r, err := http.NewRequest(http.MethodDelete, "/author/"+testCase.id, nil)
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
