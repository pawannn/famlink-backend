package domain

type UserSchema struct {
	ID      string `json:"id"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Country string `json:"country"`
}

type VerifyPhonePayload struct {
	Phone   string `json:"phone"`
	Country string `json:"country"`
	OTP     int    `json:"otp"`
}

type VerifyPhoneResponse struct {
	User  UserSchema `json:"user"`
	Token string     `json:"token"`
}
