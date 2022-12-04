package utils

import (
	"errors"
	"unicode"
)

func ValidateMinSize(password string, ruleValue int) error {
	var err error

	if len(password) < ruleValue {
		err = errors.New(MinSize)
	}

	return err
}

func ValidateMinLowercase(password string, ruleValue int) error {
	var minLowercaseErr int
	var err error

	for _, r := range password {
		if unicode.IsLower(r) {
			minLowercaseErr++
		}
	}

	if minLowercaseErr < ruleValue {
		err = errors.New(MinLowercase)
	}

	return err
}

func ValidateMinUppercase(password string, ruleValue int) error {
	var minUppercaseErr int
	var err error

	for _, r := range password {
		if unicode.IsUpper(r) {
			minUppercaseErr++
		}
	}

	if minUppercaseErr < ruleValue {
		err = errors.New(MinUppercase)
	}

	return err
}

func ValidateMinDigit(password string, ruleValue int) error{
	var minDigitOk int
	var err error

	for _, r := range password {
		if (r > 47) && (r < 58) {
			minDigitOk++
		}
	}

	if minDigitOk < ruleValue {
		err = errors.New(MinDigit)
	}

	return err
}

func ValidateMinSpecialChars(password string, ruleValue int) error{
	var specialCharsOk int
	var err error

	for _, r := range password {
		if ((r > 31) && (r < 48)) || ((r > 57) && (r < 65)) || ((r > 90) && (r < 97)) || ((r > 122) && (r < 126)) {
			specialCharsOk++
		}
	}

	if specialCharsOk < ruleValue {
		err = errors.New(MinSpecialChars)
	}

	return err
}

func ValidateNoRepeated(password string, ruleValue int) error{
	var noRepeatedErr int
	var lastChar rune
	var err error

	for _, r := range password {
		if r == lastChar {
			noRepeatedErr++
		}
		lastChar = r
	}

	if noRepeatedErr > 0 {
		noRepeatedErr++
	}

	if noRepeatedErr > ruleValue {
		err = errors.New(NoRepeated)
	}

	return err
}
