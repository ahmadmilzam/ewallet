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

func init() {
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
		return "ACTIVE", nil
	})

	_ = faker.AddProvider("accountCOA", func(v reflect.Value) (interface{}, error) {

		return "LIABILITIES", nil
	})
}

func GenerateAccountData(a *entity.Account) error {
	err := faker.FakeData(a)
	if err != nil {
		return err
	}

	faker.ResetUnique()
	return nil
}
