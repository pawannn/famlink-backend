package port

import domain "github.com/pawannn/famlink/domain/users"

type UserRepository struct {
	Repo domain.UserService
}

func InitUserService(repo domain.UserService) *UserRepository {
	return &UserRepository{Repo: repo}
}

func (uS *UserRepository) Register(user domain.UserSchema) (*domain.UserSchema, error) {
	return uS.Repo.Register(user)
}

func (uS *UserRepository) GetUserByID(id string) (*domain.UserSchema, error) {
	return uS.Repo.GetUserByID(id)
}

func (uS *UserRepository) UpdateUser(id string, name *string, avatar *string) (*domain.UserSchema, error) {
	return uS.Repo.UpdateUser(id, name, avatar)
}

func (uS *UserRepository) GetUserByPhone(phone string) (*domain.UserSchema, error) {
	return uS.Repo.GetUserByPhone(phone)
}
