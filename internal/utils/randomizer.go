package utils

import (
	// "math/rand"
	// "time"

	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/dongri/phonenumber"
	"github.com/go-faker/faker/v4"
)

/*
if !array.Contains(GetSupportedUserType(), payer.User.Type) {
	return false, errors.New("unsupported payer type")
}
*/

func CustomGenerator() {
	_ = faker.AddProvider("customphone", func(v reflect.Value) (interface{}, error) {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		num := r.Intn(99999999999)

		return fmt.Sprintf("8%d", num), nil
	})

	_ = faker.AddProvider("accountRole", func(v reflect.Value) (interface{}, error) {
		ar := usecase.GetSupportedAccountRole()
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		idx := r.Intn(len(ar) - 1)

		return ar[idx], nil
	})

	_ = faker.AddProvider("accountStatus", func(v reflect.Value) (interface{}, error) {
		as := usecase.GetSupportedAccountStatus()
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		idx := r.Intn(len(as) - 1)

		return as[idx], nil
	})
}

func RandomAccountData(a entity.Account) entity.Account {
	CustomGenerator()

	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	faker.ResetUnique()
	return a
}

func ParsePhoneNumber() {
	number := phonenumber.Parse("081284026291", "ID")
	fmt.Println(number)
}
