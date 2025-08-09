package port

import tokenService "github.com/pawannn/famlink/core/services/token"

type TokenRepo struct {
	repo tokenService.TokenService
}

func InitTokenRepo(repo tokenService.TokenService) *TokenRepo {
	tR := TokenRepo{
		repo: repo,
	}
	return &tR
}

func (tR TokenRepo) GenerateUserToken(userID string) (string, error) {
	return tR.repo.GenerateToken(userID)
}

func (tR TokenRepo) ParseUserToken(token string) (string, error) {
	return tR.repo.ParseToken(token)
}
