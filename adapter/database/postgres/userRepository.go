package database

import (
	"database/sql"
	"errors"
	"fmt"
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

func (uR UserRepo) Register(user domain.UserSchema) (*domain.UserSchema, error) {
	row := uR.db.QueryRow("INSERT INTO users(id, name, phone, country) VALUES($1, $2, $3, $4) RETURNING id, name, phone, country, COALESCE(avatar, '')", user.ID, user.Name, user.Phone, user.Country)
	var UserDetails domain.UserSchema
	if err := row.Scan(&UserDetails.ID, &UserDetails.Name, &UserDetails.Phone, &UserDetails.Country, &UserDetails.Avatar); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, errors.New(constants.ERR_USER_EXIST)
		}
		return nil, err
	}
	return &UserDetails, nil
}

func (uR UserRepo) UpdateUser(id string, name *string, avatar *string) (*domain.UserSchema, error) {
	query := strings.Builder{}
	query.WriteString("UPDATE users SET ")
	params := []any{}
	setClauses := []string{}
	paramIndex := 1
	if name != nil {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", paramIndex))
		params = append(params, *name)
		paramIndex++
	}
	if avatar != nil {
		setClauses = append(setClauses, fmt.Sprintf("avatar = $%d", paramIndex))
		params = append(params, *avatar)
		paramIndex++
	}
	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}
	query.WriteString(strings.Join(setClauses, ", "))
	query.WriteString(fmt.Sprintf(" WHERE id = $%d RETURNING id, name, phone, country, COALESCE(avatar, '')", paramIndex))
	params = append(params, id)
	row := uR.db.QueryRow(query.String(), params...)

	var user domain.UserSchema
	err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Country, &user.Avatar)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (uR UserRepo) GetUserByID(id string) (*domain.UserSchema, error) {
	row := uR.db.QueryRow("SELECT id, name, phone, country, COALESCE(avatar, '') FROM users WHERE id = $1", id)
	var UserDetails domain.UserSchema
	if err := row.Scan(&UserDetails.ID, &UserDetails.Name, &UserDetails.Phone, &UserDetails.Country, &UserDetails.Avatar); err != nil {
		if err.Error() == constants.ERR_NO_ROWS {
			return nil, nil
		}
		return nil, err
	}
	return &UserDetails, nil
}

func (uR UserRepo) GetUserByPhone(phone string) (*domain.UserSchema, error) {
	row := uR.db.QueryRow("SELECT id, name, phone, country, COALESCE(avatar, '') FROM users WHERE phone = $1", phone)
	var UserDetails domain.UserSchema
	if err := row.Scan(&UserDetails.ID, &UserDetails.Name, &UserDetails.Phone, &UserDetails.Country, &UserDetails.Avatar); err != nil {
		if err.Error() == constants.ERR_NO_ROWS {
			return nil, nil
		}
		return nil, err
	}
	return &UserDetails, nil
}
