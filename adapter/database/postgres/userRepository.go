package database

import (
	"database/sql"
	"errors"
	"strings"

	domain "github.com/pawannn/famlink/domain/users"
	"github.com/pawannn/famlink/pkg/constants"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserService {
	return UserRepo{db: db}
}

func (uR UserRepo) Register(name string, phone string, country string) (*domain.UserSchema, error) {
	row := uR.db.QueryRow("INSERT INTO users(name, phone, country) VALUES($1, $2, $3) RETURNING id, name, phone, country, COALESCE(avatar, '')", name, phone, country)
	var UserDetails domain.UserSchema
	if err := row.Scan(&UserDetails.ID, &UserDetails.Name, &UserDetails.Phone, &UserDetails.Country, &UserDetails.Avatar); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, errors.New(constants.ERR_USER_EXIST)
		}
		return nil, err
	}
	return &UserDetails, nil
}

func (uR UserRepo) GetUser(phone string) (*domain.UserSchema, error) {
	row := uR.db.QueryRow("SELECT id, name, phone, country, COALESCE(avatar, '') FROM users WHERE phone = $1", phone)
	var UserDetails domain.UserSchema
	if err := row.Scan(&UserDetails.ID, &UserDetails.Name, &UserDetails.Phone, &UserDetails.Country, &UserDetails.Avatar); err != nil {
		return nil, err
	}
	return &UserDetails, nil
}

func (uR UserRepo) UpdateUser(user *domain.UserSchema) (*domain.UserSchema, error) {
	return nil, nil
}
