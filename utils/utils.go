package utils

import (
	"errors"
	"unicode"
)

// check if the total of caracters is valid based on rule value
func ValidateMinSize(password string, ruleValue int) error {
	var err error

	if len(password) < ruleValue {
		err = errors.New(MinSize)
	}

	return err
}

// check if the total of lowercase caracters is valid based on rule value
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

// check if the total of uppercase caracters is valid based on rule value
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

// check if the total of number caracters is valid based on rule value
func ValidateMinDigit(password string, ruleValue int) error {
	var minDigitOk int
	var err error

	for _, r := range password {
		if unicode.IsNumber(r) {
			minDigitOk++
		}
	}

	if minDigitOk < ruleValue {
		err = errors.New(MinDigit)
	}

	return err
}

// check if the total of special caracters is valid based on rule value
func ValidateMinSpecialChars(password string, ruleValue int) error {
	var specialCharsOk int
	var err error

	for _, r := range password {
		if err = isSpecialRune(r); err == nil {
			specialCharsOk++
		}
	}

	if specialCharsOk < ruleValue {
		err = errors.New(MinSpecialChars)
	}

	return err
}

// check if the total of repeated caracters is valid based on rule value
func ValidateNoRepeated(password string, ruleValue int) error {
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

func isSpecialRune(r rune) error {
	if ((r > 31) && (r < 48)) || ((r > 57) && (r < 65)) || ((r > 90) && (r < 97)) || ((r > 122) && (r < 126)) {
		return nil
	}

	return errors.New(SpecialCharacterNotFound)
}
