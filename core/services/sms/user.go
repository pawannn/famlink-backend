package sms

type UserSmsService interface {
	SendUserOTP(phoneNumber string) error
	VerifyUserOTP(phoneNumber string, OTP string) (bool, error)
}
