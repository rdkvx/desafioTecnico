package services

import (
	"github.com/rdkvx/desafioTecnico/v2/models"
	"github.com/rdkvx/desafioTecnico/v2/utils"
)

type IpasswordValidator interface{
	ValidatePassword(bodyRequest models.ValidatePassword) ([]string)
}

type Validator struct{}

var PasswordValidator IpasswordValidator

func NewPasswordValidator() IpasswordValidator{
	return &Validator{}
}

func (vp *Validator) ValidatePassword(bodyRequest models.ValidatePassword) ([]string) {

	failedRules := []string{}

	// validate every rule that came from the request
	for _, rule := range bodyRequest.Rules {
		switch rule.Rule {
		case utils.MinSize:
			if err := utils.ValidateMinSize(bodyRequest.Password, rule.Value); err != nil {
				failedRules = append(failedRules, utils.MinSize)
			}

		case utils.MinLowercase:
			if err := utils.ValidateMinLowercase(bodyRequest.Password, rule.Value); err != nil {
				failedRules = append(failedRules, utils.MinLowercase)
			}
		
		case utils.MinUppercase:
			if err := utils.ValidateMinUppercase(bodyRequest.Password, rule.Value); err != nil {
				failedRules = append(failedRules, utils.MinUppercase)
			}

		case utils.MinDigit:
			if err := utils.ValidateMinDigit(bodyRequest.Password, rule.Value); err != nil {
				failedRules = append(failedRules, utils.MinDigit)
			}

		case utils.MinSpecialChars:
			if err := utils.ValidateMinSpecialChars(bodyRequest.Password, rule.Value); err != nil {
				failedRules = append(failedRules, utils.MinSpecialChars)
			}

		case utils.NoRepeated:
			if err := utils.ValidateNoRepeated(bodyRequest.Password, rule.Value); err != nil {
				failedRules = append(failedRules, utils.NoRepeated)
			}
		}
	}

	return failedRules
}
