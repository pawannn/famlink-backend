package helpers

import "strings"

var validISOCountries = map[string]struct{}{
	"US": {}, "IN": {}, "GB": {}, "CA": {}, "AU": {}, "DE": {}, "FR": {}, "CN": {}, "JP": {},
}

func ValidateCountry(code string) bool {
	_, exists := validISOCountries[strings.ToUpper(code)]
	return exists
}
