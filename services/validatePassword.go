package services

import (
	"regexp"
	"unicode"

	"github.com/rdkvx/desafioTecnico/v2/model"
	//"github.com/rdkvx/desafioTecnico/v2/utils"
)

//func ValidatePassword(bodyRequest map[string]interface{}) error {
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
			minDigitErr := 0

			re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
			subMatchAll := re.FindAllString(bodyRequest.Password, -1)
			for range subMatchAll {
				minDigitErr++
			}

			if minDigitErr < rule.Value {
				failedRules = append(failedRules, "minDigit")
			}

		case "minSpecialChars":
			specialCharsErr := 0
			re := regexp.MustCompile(`[^a-zA-Z0-9]+`)

			subMatchAll := re.FindAllString(bodyRequest.Password, -1)
			for range subMatchAll {
				specialCharsErr++
			}

			if specialCharsErr < rule.Value {
				failedRules = append(failedRules, "minSpecialChars")
			}

		case "noRepeated":
			noRepeatedErr := 0
			var lastChar rune
			for _, r := range bodyRequest.Password{
				if r == lastChar{
					noRepeatedErr++
				}
				lastChar = r
			}

			if noRepeatedErr > 0{
				noRepeatedErr++
			}

			if noRepeatedErr > rule.Value{
				failedRules = append(failedRules, "noRepeated")
			}
		}

	}

	return
}

/* func validateRules(rules []interface{}) error {

	for _, rule := range rules {
		if
	}

	return nil
} */
