package services

import "github.com/pawannn/famly/internal/core/domain"

type SmsManager struct {
	repo domain.SmsRepo
}

func InitSmsManager(repo domain.SmsRepo) SmsManager {
	return SmsManager{
		repo: repo,
	}
}

func (sM SmsManager) SendOTP(phoneNumber string) error {
	return sM.repo.SendOTP(phoneNumber)
}

func (sM SmsManager) VerifyOTP(phoneNumber string, OTP string) (bool, error) {
	return sM.repo.VerifyOTP(phoneNumber, OTP)
}
