package services

import "github.com/pawannn/famly/internal/core/domain"

type AuthManager struct {
	repo domain.AuthorizationRepo
}

func InitAuthManager(repo domain.AuthorizationRepo) *AuthManager {
	return &AuthManager{
		repo: repo,
	}
}

func (aM *AuthManager) GenerateToken(userID string) (string, error) {
	return aM.repo.GenerateToken(userID)
}

func (aM *AuthManager) ParseToken(token string) (string, error) {
	return aM.repo.ParseToken(token)
}
