package user

import datastoredomain "github.com/pawannn/famly/internal/core/domain/datastore"

type PhoneVerificationPayload struct {
	Country string `json:"country"`
	Phone   string `json:"phone"`
	OTP     string `json:"otp"`
}

type PhoneVerificationResponse struct {
	User  *datastoredomain.UserSchema
	Token string
}
