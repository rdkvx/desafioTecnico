package services

import (
	"unicode"

	"github.com/rdkvx/desafioTecnico/v2/model"
)

func ValidatePassword(bodyRequest model.ValidatePassword) (failedRules []string) {

	for _, rule := range bodyRequest.Rules {
		switch rule.Rule {
		case "minSize":
			if len(bodyRequest.Password) < rule.Value {
				failedRules = append(failedRules, "minSize")
			}

		case "minLowercase":
			minLowercaseErr := 0

			for _, r := range bodyRequest.Password {
				if unicode.IsLower(r) {
					minLowercaseErr++
				}
			}

			if minLowercaseErr < rule.Value {
				failedRules = append(failedRules, "minLowercase")
			}

		case "minUppercase":
			minUppercaseErr := 0

			for _, r := range bodyRequest.Password {
				if unicode.IsUpper(r) {
					minUppercaseErr++
				}
			}

			if minUppercaseErr < rule.Value {
				failedRules = append(failedRules, "minUppercase")
			}

		case "minDigit":
			minDigitOk := 0

			for _, r := range bodyRequest.Password {
				if (r > 47) && (r < 58) {
					minDigitOk++
				}
			}

			if minDigitOk < rule.Value {
				failedRules = append(failedRules, "minDigit")
			}

		case "minSpecialChars":
			specialCharsOk := 0
			
			for _, r := range bodyRequest.Password {
				if ((r > 31) && (r < 48)) || ((r > 57) && (r < 65)) || ((r > 90) && (r < 97)) || ((r > 122) && (r < 126)) {
					specialCharsOk++
				}
			}

			if specialCharsOk < rule.Value {
				failedRules = append(failedRules, "minSpecialChars")
			}

		case "noRepeated":
			noRepeatedErr := 0
			var lastChar rune
			for _, r := range bodyRequest.Password {
				if r == lastChar {
					noRepeatedErr++
				}
				lastChar = r
			}

			if noRepeatedErr > 0 {
				noRepeatedErr++
			}

			if noRepeatedErr > rule.Value {
				failedRules = append(failedRules, "noRepeated")
			}
		}
	}

	return
}
