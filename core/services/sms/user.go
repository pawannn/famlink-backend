package sms

type UserSmsService interface {
	SendUserOTP(phoneNumber string) error
}
