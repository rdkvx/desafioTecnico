package services

import (
	"github.com/rdkvx/desafioTecnico/v2/model"
	//"github.com/rdkvx/desafioTecnico/v2/utils"
)

//func ValidatePassword(bodyRequest map[string]interface{}) error {
func ValidatePassword(bodyRequest model.ValidatePassword) error {
	/*
		password, ok := bodyRequest["password"].(string)
		if !ok || password == "" {
			return fmt.Errorf("failed to get password")
		}

		rules, ok := bodyRequest["rules"].([]interface{})
		if !ok || len(rules) == 0 {
			return fmt.Errorf("failed to get rules")
		} */

	return nil
}

/* func validateRules(rules []interface{}) error {

	for _, rule := range rules {
		if
	}

	return nil
} */
