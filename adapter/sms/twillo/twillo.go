package sms

import (
	"fmt"

	"github.com/pawannn/famly/core/services/sms"
	appconfig "github.com/pawannn/famly/pkg/appConfig"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilloRepo struct {
	Client    *twilio.RestClient
	ServiceID string
}

func InitTwilloClient(c appconfig.Config) sms.SmsService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.SMS_Account_Sid,
		Password: c.SMS_Service_Token,
	})
	return TwilloRepo{
		Client:    client,
		ServiceID: c.SMS_Service_ID,
	}
}

func (tR TwilloRepo) SendOTP(phoneNumber string) error {
	params := &verify.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := tR.Client.VerifyV2.CreateVerification(tR.ServiceID, params)
	if err != nil {
		return fmt.Errorf("failed to send OTP: %w", err)
	}

	if resp != nil && resp.Sid != nil {
		fmt.Printf("OTP sent. SID: %s, Status: %s\n", *resp.Sid, *resp.Status)
	}

	return nil
}

func (tR TwilloRepo) VerifyOTP(phoneNumber string, code string) (bool, error) {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := tR.Client.VerifyV2.CreateVerificationCheck(tR.ServiceID, params)
	if err != nil {
		return false, fmt.Errorf("failed to verify OTP: %w", err)
	}

	if resp != nil && resp.Status != nil {
		if *resp.Status == "approved" {
			return true, nil
		}
		return false, nil
	}

	return false, fmt.Errorf("invalid response from Twilio")
}
