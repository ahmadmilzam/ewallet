package utils

import (
	"fmt"
	"testing"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/dongri/phonenumber"
	"github.com/go-faker/faker/v4"
	// "github.com/stretchr/testify/assert"
)

func TestRandomAccount(t *testing.T) {
	// assert := assert.New(t)

	r := RandomAccountData(entity.Account{})
	fmt.Printf("%+v", r)

	// assert equality
	// assert.Equal(123, 123, "they should be equal")

	// assert inequality
	// assert.NotEqual(123, 456, "they should not be equal")
}

func TestPhone(t *testing.T) {
	CustomGenerator()
	var a entity.Account
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	faker.ResetUnique()
	fmt.Println(a.Phone)
	a.Phone = phonenumber.Parse(a.Phone, "ID")
	fmt.Println(a.Phone)
}
