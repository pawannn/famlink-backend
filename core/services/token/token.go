package services

type TokenService interface {
	GenerateToken(userID string) (string, error)
	ParseToken(token string) (string, error)
}
