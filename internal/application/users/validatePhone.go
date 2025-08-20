package users

import (
	"net/http"
	"strings"

	"github.com/nyaruka/phonenumbers"
	"github.com/pawannn/famly/internal/pkg/constants"
	"github.com/pawannn/famly/internal/pkg/helpers"
)

func (uA UserApplication) ValidatePhone(country string, phone string) helpers.FamlyErr {
	countryCode := strings.ToUpper(strings.TrimSpace(country))
	if !helpers.ValidateCountry(countryCode) {
		return helpers.FamlyErr{
			Code:    http.StatusBadRequest,
			Message: constants.ERR_INVALID_COUNTRY_CODE,
			Error:   nil,
		}
	}

	parsedPhone, err := phonenumbers.Parse(phone, countryCode)
	if err != nil || !phonenumbers.IsValidNumber(parsedPhone) {
		return helpers.FamlyErr{
			Code:    http.StatusBadRequest,
			Message: constants.ERR_INVALID_PHONE,
			Error:   err,
		}
	}

	formattedPhone := phonenumbers.Format(parsedPhone, phonenumbers.E164)
	err = uA.SmsRepo.SendOTP(formattedPhone)
	if err != nil {
		return helpers.FamlyErr{
			Code:    http.StatusInternalServerError,
			Message: "Unable to send OTP",
			Error:   err,
		}
	}
	return helpers.NoFamErr
}
