package port

import tokenService "github.com/pawannn/famlink/core/services"

type TokenRepo struct {
	repo tokenService.TokenService
}

func InitTokenRepo(repo tokenService.TokenService) *TokenRepo {
	tR := TokenRepo{
		repo: repo,
	}
	return &tR
}

func (tR TokenRepo) GenerateToken(userID string) (string, error) {
	return tR.repo.GenerateToken(userID)
}

func (tR TokenRepo) ParseToken(token string) (string, error) {
	return tR.repo.ParseToken(token)
}
