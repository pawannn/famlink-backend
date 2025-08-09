package database

import domain "github.com/pawannn/famlink/core/domain/users"

type UserDBport struct {
	Repo domain.UserService
}

func InitUserDBPort(repo domain.UserService) *UserDBport {
	return &UserDBport{Repo: repo}
}

func (uS *UserDBport) Register(user domain.UserSchema) (*domain.UserSchema, error) {
	return uS.Repo.Register(user)
}

func (uS *UserDBport) GetUserByID(id string) (*domain.UserSchema, error) {
	return uS.Repo.GetUserByID(id)
}

func (uS *UserDBport) UpdateUser(id string, name *string, avatar *string) (*domain.UserSchema, error) {
	return uS.Repo.UpdateUser(id, name, avatar)
}

func (uS *UserDBport) GetUserByPhone(phone string) (*domain.UserSchema, error) {
	return uS.Repo.GetUserByPhone(phone)
}
