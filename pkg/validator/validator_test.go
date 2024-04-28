package validator_test

import (
	"testing"

	"github.com/ahmadmilzam/ewallet/pkg/validator"
	"github.com/stretchr/testify/assert"
)

func TestIsValidAmount(t *testing.T) {
	invalidParam := int64(-1999)
	validParam := int64(1)

	assert.Equal(t, false, validator.IsValidAmount(invalidParam))
	assert.Equal(t, true, validator.IsValidAmount(validParam))
}

func TestValide164Phone(t *testing.T) {
	testCases := []struct {
		desc    string
		param   string
		isValid bool
	}{
		{
			desc:    "valid e164 phone format",
			param:   "+5111111111",
			isValid: true,
		},
		{
			desc:    "phone doesn't have a plus sign",
			param:   "5111111111",
			isValid: false,
		},
		{
			desc:    "phone must be numeric",
			param:   "+511111asdf",
			isValid: false,
		},
		{
			desc:    "phone total numeric length less than min char (8)",
			param:   "+511111",
			isValid: false,
		},
		{
			desc:    "phone total numeric length excced max char (15)",
			param:   "+5111111111111111",
			isValid: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.isValid, validator.IsValidPhone(tC.param))
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		desc    string
		param   string
		isValid bool
	}{
		{
			desc:    "valid common email format",
			param:   "owner@gmail.com",
			isValid: true,
		},
		{
			desc:    "email has no '@'",
			param:   "asd.gmail.com",
			isValid: false,
		},
		{
			desc:    "email has no '.' to indicate a domain name",
			param:   "asd.gmail_com",
			isValid: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.isValid, validator.IsValidEmail(tC.param))
		})
	}
}

func TestIsValidTimestampt(t *testing.T) {
	validTime := "2024-04-16T22:58:50+07:00"
	noTz := "2024-04-16T22:58:50"
	noSeparator := "2024-04-16 22:58:50+07:00"
	noSeparatorAndTz := "2024-04-16 22:58:50"
	random := "26/06/1900"

	assert.Equal(t, true, validator.IsValidTimestampt(validTime))
	assert.Equal(t, false, validator.IsValidTimestampt(noTz))
	assert.Equal(t, false, validator.IsValidTimestampt(noSeparator))
	assert.Equal(t, false, validator.IsValidTimestampt(noSeparatorAndTz))
	assert.Equal(t, false, validator.IsValidTimestampt(random))
}
