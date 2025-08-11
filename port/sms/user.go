package sms

import "github.com/pawannn/famly/core/services/sms"

type UserSmsPort struct {
	Repo sms.UserSmsService
}

func InitUserSmsPort(repo sms.UserSmsService) UserSmsPort {
	return UserSmsPort{
		Repo: repo,
	}
}

func (uS UserSmsPort) SendUserOTP(phoneNumber string) error {
	return uS.Repo.SendUserOTP(phoneNumber)
}

func (uS UserSmsPort) VerifyUserOTP(phoneNumber string, OTP string) (bool, error) {
	return uS.Repo.VerifyUserOTP(phoneNumber, OTP)
}
