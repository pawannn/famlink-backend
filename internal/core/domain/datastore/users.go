package datastoredomain

type UserSchema struct {
	ID      string `json:"id"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Country string `json:"country"`
}

type UserDBRepo interface {
	Register(UserSchema) (*UserSchema, error)
	GetUserByID(id string) (*UserSchema, error)
	UpdateUser(id string, name *string, avatar *string) (*UserSchema, error)
	GetUserByPhone(phone string) (*UserSchema, error)
}
