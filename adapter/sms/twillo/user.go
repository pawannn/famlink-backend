package sms

import "github.com/pawannn/famly/core/services/sms"

type UserSmsRepo struct {
	SmsClient sms.SmsService
}

func InitUserSmsRepo(smsClient sms.SmsService) sms.UserSmsService {
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
