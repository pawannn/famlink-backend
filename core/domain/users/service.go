package domain

type UserService interface {
	Register(UserSchema) (*UserSchema, error)
	GetUserByID(id string) (*UserSchema, error)
	UpdateUser(id string, name *string, avatar *string) (*UserSchema, error)
	GetUserByPhone(phone string) (*UserSchema, error)
}
