package port

import domain "github.com/pawannn/famlink/domain/users"

type UserRepository struct {
	Repo domain.UserService
}

func InitUserService(repo domain.UserService) *UserRepository {
	return &UserRepository{Repo: repo}
}

func (uS *UserRepository) Register(name string, phone string, country string) (*domain.UserSchema, error) {
	return uS.Repo.Register(name, phone, country)
}

func (uS *UserRepository) GetUser(phone string) (*domain.UserSchema, error) {
	return uS.Repo.GetUser(phone)
}

func (uS *UserRepository) UpdateUser(user *domain.UserSchema) (*domain.UserSchema, error) {
	return uS.Repo.UpdateUser(user)
}
