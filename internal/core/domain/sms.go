package domain

type SmsRepo interface {
	SendOTP(phoneNumber string) error
	VerifyOTP(phoneNumber string, OTP string) (bool, error)
}
