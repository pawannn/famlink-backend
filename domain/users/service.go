package domain

type UserService interface {
	Register(name string, phone string, country string) (*UserSchema, error)
	GetUser(phone string) (*UserSchema, error)
	UpdateUser(user *UserSchema) (*UserSchema, error)
}
