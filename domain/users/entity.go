package domain

type UserSchema struct {
	ID      string `json:"id"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Country string `json:"country"`
}
