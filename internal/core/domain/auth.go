package domain

type AuthorizationRepo interface {
	GenerateToken(userID string) (string, error)
	ParseToken(token string) (string, error)
}
