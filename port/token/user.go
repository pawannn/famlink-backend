package token

import tokenService "github.com/pawannn/famlink/core/services/token"

type TokenPort struct {
	repo tokenService.TokenService
}

func InitTokenPort(repo tokenService.TokenService) *TokenPort {
	tR := TokenPort{
		repo: repo,
	}
	return &tR
}

func (tR TokenPort) GenerateUserToken(userID string) (string, error) {
	return tR.repo.GenerateToken(userID)
}

func (tR TokenPort) ParseUserToken(token string) (string, error) {
	return tR.repo.ParseToken(token)
}
