package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	appconfig "github.com/pawannn/famlink/pkg/appConfig"
)

type TokenRepo struct {
	Secret string
}

type claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func InitTokenService(c appconfig.Config) TokenRepo {
	return TokenRepo{
		Secret: c.Token_secret,
	}
}

func (tr TokenRepo) GenerateJWT(userID string) (string, error) {
	claims := claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tr.Secret))
}

func (tr TokenRepo) ValidateJWT(tokenStr string) (*claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tr.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid or expired token")
}
