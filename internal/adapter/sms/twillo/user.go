package sms

import "github.com/pawannn/famly/internal/core/domain"

type UserSmsRepo struct {
	SmsClient domain.SmsRepo
}

func InitUserSmsRepo(smsClient domain.SmsRepo) UserSmsRepo {
	return UserSmsRepo{
		SmsClient: smsClient,
	}
}

func (uS UserSmsRepo) SendUserOTP(phone string) error {
	err := uS.SmsClient.SendOTP(phone)
	return err
}

func (uS UserSmsRepo) VerifyUserOTP(phoneNumber string, code string) (bool, error) {
	ok, err := uS.SmsClient.VerifyOTP(phoneNumber, code)
	return ok, err
}
