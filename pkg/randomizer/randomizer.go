package randomizer

import (
	// "math/rand"
	// "time"

	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
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

		return fmt.Sprintf("62%d", num), nil
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

	_ = faker.AddProvider("accountCOA", func(v reflect.Value) (interface{}, error) {

		return "LIABILITIES", nil
	})
}

func RandomAccountData(a *entity.Account) error {
	CustomGenerator()

	err := faker.FakeData(a)
	if err != nil {
		return err
	}
	faker.ResetUnique()
	return nil
}
