package sms

type SmsService interface {
	SendOTP(phoneNumber string) error
	VerifyOTP(phoneNumber string, OTP string) (bool, error)
}
