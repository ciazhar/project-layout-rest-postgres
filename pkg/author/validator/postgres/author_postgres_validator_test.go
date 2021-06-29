package postgres

//func init() {
//	env.InitPath(env.GetEnvPath() + "/config.json")
//	query.PG()
//	validator.Init()
//}
//
//func NewRole() model.Role {
//	var role model.Role
//	testdata.ToStruct("role/actual.1.golden", &role)
//	repo := postgres.NewRolePostgresRepository()
//	repo.Store(&role)
//	return role
//}
//
//type RoleExample struct {
//	RoleId string `validate:"roleMustExist"`
//}
//
//func NewPostExample() RoleExample {
//	return RoleExample{RoleId: NewRole().ID.String()}
//}
//
//func TestRolePostgresValidatorInit(t *testing.T) {
//
//	t.Run("default", func(t *testing.T) {
//		dummy := NewPostExample()
//		repo := new(mocks.RolePostgresRepository)
//		repo.On("GetByID", dummy.RoleId).Return(NewRole(), nil)
//		NewRolePostgresValidator(repo)
//
//		err := validator.Struct(dummy)
//		assert.NoError(t, err)
//	})
//
//	t.Run("error", func(t *testing.T) {
//		dummy := NewPostExample()
//		repo := new(mocks.RolePostgresRepository)
//		repo.On("GetByID", dummy.RoleId).Return(NewRole(), errors.New("not found"))
//		NewRolePostgresValidator(repo)
//
//		err := validator.Struct(dummy)
//		assert.Error(t, err)
//	})
//
//}
