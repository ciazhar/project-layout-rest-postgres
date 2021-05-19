package usecase

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
//func TestAuthorUseCase_Store(t *testing.T) {
//	repo := new(mocks.AuthorPostgresRepository)
//	uc := NewAuthorUseCase(repo)
//	testCases := []struct {
//		name        string
//		author model.Author
//		returnError error
//	}{
//		{"default", NewActual(), nil},
//		{"default2", NewActual2(), nil},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			repo.On("Store", &testCase.author).Return(testCase.returnError)
//
//			err := uc.Store(&testCase.author)
//
//			assert.NoError(t, err)
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestAuthorUseCase_Fetch(t *testing.T) {
//	repo := new(mocks.AuthorPostgresRepository)
//	uc := NewAuthorUseCase(repo)
//	testCases := []struct {
//		name             string
//		offset           int
//		limit            int
//		returnAuthor []model.Author
//		returnError      error
//	}{
//		{"default", 0, 10, []model.Author{NewActual(), NewActual2()}, nil},
//		{"default2", 0, 5, []model.Author{NewActual(), NewActual2()}, nil},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//
//			param := rest.NewParam()
//			param.Offset = 1
//			param.Limit = 10
//
//			repo.On("Fetch", param).Return(testCase.returnAuthor, testCase.returnError)
//
//			expected, err := uc.Fetch(param)
//
//			assert.NotEmpty(t, expected)
//			assert.NoError(t, err)
//			assert.Len(t, expected, len(testCase.returnAuthor))
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestAuthorUseCase_GetByID(t *testing.T) {
//	repo := new(mocks.AuthorPostgresRepository)
//	uc := NewAuthorUseCase(repo)
//	testCases := []struct {
//		name             string
//		id               string
//		returnAuthor model.Author
//		returnError      error
//	}{
//		{"default", "1", NewActual(), nil},
//		{"default2", "2", NewActual(), nil},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			repo.On("GetByID", testCase.id).Return(testCase.returnAuthor, testCase.returnError)
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
//func TestAuthorUseCase_Update(t *testing.T) {
//	repo := new(mocks.AuthorPostgresRepository)
//	uc := NewAuthorUseCase(repo)
//	actual := NewActual()
//	actual.Id = "100"
//	testCases := []struct {
//		name        string
//		author model.Author
//		returnError error
//	}{
//		{"default", NewActual(), nil},
//		{"default2", actual, errors.New("not found")},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.name, func(t *testing.T) {
//			repo.On("GetByID", testCase.author.Id).Return(testCase.author, testCase.returnError)
//			repo.On("Update", &testCase.author).Return(testCase.returnError)
//
//			err := uc.Update(&testCase.author)
//
//			assert.Equal(t, err, testCase.returnError)
//			repo.AssertExpectations(t)
//		})
//	}
//}
//
//func TestAuthorUseCase_Delete(t *testing.T) {
//	actual := NewActual()
//	actual.Id = "100"
//	repo := new(mocks.AuthorPostgresRepository)
//	uc := NewAuthorUseCase(repo)
//	testCases := []struct {
//		name        string
//		author model.Author
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
//			repo.On("GetByID", testCase.author.Id).Return(testCase.author, testCase.returnError)
//
//			testCase.author.DeletedAt = time.Now()
//			repo.On("Update", &testCase.author).Return(testCase.returnError)
//
//			err := uc.Delete(testCase.author.Id)
//
//			assert.Equal(t, err, testCase.returnError)
//			//repo.AssertExpectations(t)
//		})
//	}
//}
