package services

import (
	"fmt"
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
			hasLower := 0

			for _, r := range bodyRequest.Password {
				if unicode.IsLower(r) {
					hasLower++
				}
			}

			if hasLower < rule.Value {
				failedRules = append(failedRules, "minSize")
			}

		case "minSpecialChars":
			hasSpecialChars := 0
			re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
			
			submatchall := re.FindAllString(bodyRequest.Password, -1)
			for _, element := range submatchall{
				fmt.Println(element)
				hasSpecialChars++
			}

			if hasSpecialChars < rule.Value{
				failedRules = append(failedRules, "minSpecialChars")
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
