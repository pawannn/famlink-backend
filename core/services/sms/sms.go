package sms

type SmsService interface {
	SendOTP(phoneNumber string) error
}
